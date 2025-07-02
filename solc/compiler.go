package solc

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/compiler"
)

var versionRegexp = regexp.MustCompile(`([0-9]+)\.([0-9]+)\.([0-9]+)`)


// Solidity contains information about the solidity compiler.
type Solidity struct {
	Path, Version, FullVersion string
	Major, Minor, Patch        int
}


func (s *Solidity) makeArgs() []string {
	p := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc",
		"--optimize",                  // code optimizer switched on
		"--allow-paths", "., ./, ../", // default to support relative paths
	}
	if s.Major > 0 || s.Minor > 4 || s.Patch > 6 {
		p[1] += ",metadata,hashes"
	}
	return p
}

// SolidityVersion runs solc and parses its version output.
func SolidityVersion(solc string) (*Solidity, error) {
	if solc == "" {
		solc = "solc"
	}
	var out bytes.Buffer
	cmd := exec.Command(solc, "--version")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	matches := versionRegexp.FindStringSubmatch(out.String())
	if len(matches) != 4 {
		return nil, fmt.Errorf("can't parse solc version %q", out.String())
	}
	s := &Solidity{Path: cmd.Path, FullVersion: out.String(), Version: matches[0]}
	if s.Major, err = strconv.Atoi(matches[1]); err != nil {
		return nil, err
	}
	if s.Minor, err = strconv.Atoi(matches[2]); err != nil {
		return nil, err
	}
	if s.Patch, err = strconv.Atoi(matches[3]); err != nil {
		return nil, err
	}
	return s, nil
}

// compileSolidityString builds and returns all the contracts contained within a source string.
func compileSolidityString(solc, source string) (map[string]*compiler.Contract, error) {
	if len(source) == 0 {
		return nil, errors.New("solc: empty source string")
	}
	s, err := SolidityVersion(solc)
	if err != nil {
		return nil, err
	}
	args := append(s.makeArgs(), "--")
	cmd := exec.Command(s.Path, append(args, "-")...)
	cmd.Stdin = strings.NewReader(source)
	return s.run(cmd, source)
}

func (s *Solidity) run(cmd *exec.Cmd, source string) (map[string]*compiler.Contract, error) {
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %v\n%s", err, stderr.Bytes())
	}
	return compiler.ParseCombinedJSON(stdout.Bytes(), source, s.Version, s.Version, strings.Join(s.makeArgs(), " "))
}

