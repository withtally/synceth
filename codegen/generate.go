package codegen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type InputType struct {
	Name  string
	Alias *string
	Type  interface{}
}

type HandlersConfig struct {
	Generate   bool
	InputTypes []InputType
}

type SetupConfig struct {
	InputTypes []InputType
}

type BindingsConfig struct {
	Fakes    bool
	Handlers HandlersConfig
	Setup    SetupConfig
}

func GenerateBindings(path string, outdir string, config *BindingsConfig) error {
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		return err
	}

	if err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			if matched, err := filepath.Match("*.abi.json", filepath.Base(path)); err != nil {
				return err
			} else if matched {
				dir, fn := filepath.Split(path)
				typ := strings.TrimSuffix(fn, ".abi.json")
				name := strings.ToLower(typ)
				pkg := strings.ToLower(filepath.Base(outdir))

				if outdir == "" {
					pkg = filepath.Base(dir)
					outdir = dir
				}

				abi, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("reading abi: %w", err)
				}

				bin, err := os.ReadFile(filepath.Join(dir, typ+".bin"))
				if errors.Is(err, os.ErrNotExist) {
					bin = []byte{}
				} else if err != nil {
					return fmt.Errorf("reading abi: %w", err)
				}

				abis := []string{string(abi)}
				types := []string{typ}
				bins := []string{string(bin)}

				// We generate handlers first to avoid generating them for
				// fakes created below.
				if config.Handlers.Generate {
					handler, err := GenerateProcessor(types, abis, pkg, config.Handlers.InputTypes, config.Setup.InputTypes)
					if errors.Is(err, ErrNoEvents) {
						return nil
					} else if err != nil {
						return fmt.Errorf("generating handler: %w", err)
					}

					if err := os.WriteFile(filepath.Join(outdir, name+".handlers.go"), []byte(handler), 0600); err != nil {
						return fmt.Errorf("writing handler: %w", err)
					}
				}

				if config.Fakes {
					fake, err := GenerateFake(types[0], abis[0], pkg)
					if err != nil {
						return fmt.Errorf("generating fake: %w", err)
					}

					md, err := ParseContract(string(fake))
					if err != nil {
						return err
					}

					types = append(types, md.Types...)
					bins = append(bins, md.Bins...)
					abis = append(abis, md.ABIs...)
				}

				// Generate the contract binding
				src, err := bind.Bind(types, abis, bins, nil, pkg, bind.LangGo, nil, nil)
				if err != nil {
					return fmt.Errorf("binding abi: %w", err)
				}

				if err := os.WriteFile(filepath.Join(outdir, name+".go"), []byte(src), 0600); err != nil {
					return fmt.Errorf("writing ABI binding: %w", err)
				}
			}
			return nil
		},
	); err != nil {
		return err
	}

	return nil
}
