package solc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/shibukawa/configdir"
	"github.com/withtally/synceth/parser"
)

// See: https://docs.soliditylang.org/en/latest/installing-solidity.html#static-binaries
const solcBaseURL = "https://binaries.soliditylang.org/"

var platforms = map[string]string{
	"emscripten": "emscripten",
	"linux":      "linux",
	"darwin":     "macosx",
	"windows":    "windows",
}

var architectures = map[string]string{
	"wasm":  "wasm32",
	"amd64": "amd64",
}

var (
	cache    = configdir.New("withtally", "ethgen").QueryCacheFolder()
	buildsFn = "builds.json"
	builds   []Build
)

type Build struct {
	Path      string          `json:"path,omitempty"`
	Version   *semver.Version `json:"version,omitempty"`
	Build     string          `json:"build,omitempty"`
	Keccak256 string          `json:"keccak256,omitempty"`
	Sha256    string          `json:"sha256,omitempty"`
}

type list struct {
	Builds []Build `json:"builds,omitempty"`
}

func init() {
	b, err := cache.ReadFile(buildsFn)
	if os.IsNotExist(err) {
		return
	} else if err != nil {
		log.Fatalf("reading builds.json: %v", err)
	}

	if err := json.Unmarshal(b, &builds); err != nil {
		log.Fatalf("unmarshalling builds.json: %v", err)
	}
}

func local(vc *parser.VersionConstraint) Build {
	sort.Slice(builds, func(i, j int) bool {
		return builds[i].Version.GreaterThan(builds[j].Version)
	})

	for i := len(builds) - 1; i >= 0; i-- {
		if vc.Check(builds[i].Version) {
			return builds[i]
		}
	}

	return Build{}
}

func fetch(b Build) (Build, error) {
	uri, err := url.Parse(solcBaseURL)
	if err != nil {
		return Build{}, fmt.Errorf("parsing soliditylang url: %w", err)
	}

	log.Printf("Fetching solc compiler: %s\n", b.Path)
	uri.Path = filepath.Join(platforms[runtime.GOOS]+"-"+architectures[runtime.GOARCH], b.Path)
	res, err := http.Get(uri.String())
	if err != nil {
		return Build{}, fmt.Errorf("create solc get request: %w", err)
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return Build{}, fmt.Errorf("reding list.json: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return Build{}, fmt.Errorf("get request: %s", string(buf))
	}

	if err := cache.CreateParentDir(b.Path); err != nil {
		return Build{}, err
	}

	if err := ioutil.WriteFile(filepath.Join(cache.Path, b.Path), buf, 0555); err != nil {
		return Build{}, fmt.Errorf("writing solc: %w", err)
	}

	builds = append(builds, b)

	marshalled, err := json.Marshal(builds)
	if err != nil {
		return Build{}, fmt.Errorf("marshalling builds.json: %w", err)
	}

	if err := cache.WriteFile("builds.json", marshalled); err != nil {
		return Build{}, fmt.Errorf("writing builds.json: %w", err)
	}

	return b, nil
}

func resolve(vc *parser.VersionConstraint) (Build, error) {
	b := local(vc)
	if b.Version != nil {
		return b, nil
	}

	uri, err := url.Parse(solcBaseURL)
	if err != nil {
		return Build{}, fmt.Errorf("parsing soliditylang url: %w", err)
	}

	platform, ok := platforms[runtime.GOOS]
	if !ok {
		return Build{}, fmt.Errorf("platform not supported: %s", runtime.GOOS)
	}

	arch, ok := architectures[runtime.GOARCH]
	if !ok {
		return Build{}, fmt.Errorf("architecture not supported: %s", runtime.GOARCH)
	}

	uri.Path = filepath.Join(platform+"-"+arch, "list.json")
	res, err := http.Get(uri.String())
	if err != nil {
		return Build{}, fmt.Errorf("create solc list get request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Build{}, fmt.Errorf("reading solc list get response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return Build{}, fmt.Errorf("get request: %s", string(body))
	}

	var ls list
	if err := json.Unmarshal(body, &ls); err != nil {
		return Build{}, fmt.Errorf("unmarshaling solc list: %w", err)
	}

	sort.Slice(ls.Builds, func(i, j int) bool {
		return ls.Builds[i].Version.GreaterThan(ls.Builds[j].Version)
	})

	for _, b := range ls.Builds {
		if vc.Check(b.Version) {
			return fetch(b)
		}
	}

	return Build{}, fmt.Errorf("cant satisfy version constraint: %s", vc.String())
}

func CompileSolidityString(src string) (map[string]*compiler.Contract, error) {
	c, err := parser.NewVersionConstraint(src)
	if err != nil {
		return nil, fmt.Errorf("parsing solidity version: %w", err)
	}

	b, err := resolve(c)
	if err != nil {
		return nil, fmt.Errorf("resolving compiler: %w", err)
	}

	return compiler.CompileSolidityString(filepath.Join(cache.Path, b.Path), src)
}
