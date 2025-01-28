package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/withtally/synceth/codegen"
)

var compileCmd = &cli.Command{
	Name:  "compile",
	Usage: "compile ethereum contract bindings",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "outdir", Usage: "output directory"},
		&cli.BoolFlag{Name: "abis", Usage: "write abis"},
		&cli.BoolFlag{Name: "handlers", Usage: "write handlers"},
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
	},
	Action: func(ctx *cli.Context) error {
		path, err := filepath.Abs(ctx.Args().First())
		if err != nil {
			return err
		}

		outdir, err := filepath.Abs(ctx.String("outdir"))
		if err != nil {
			return err
		}

		if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
			return err
		}

		if err := os.Chdir(path); err != nil {
			return fmt.Errorf("changing working directory: %w", err)
		}

		if err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				if matched, err := filepath.Match("*.sol", filepath.Base(path)); err != nil {
					return err
				} else if matched {
					dir, _ := filepath.Split(path)

					if outdir == "" {
						outdir = dir
					}

					buf, err := os.ReadFile(path)
					if err != nil {
						return fmt.Errorf("reading solidity: %w", err)
					}

					md, err := codegen.ParseContract(string(buf))
					if err != nil {
						return err
					}

					for i, t := range md.Types {
						var pretty bytes.Buffer
						if err := json.Indent(&pretty, []byte(md.ABIs[i]), "", "\t"); err != nil {
							return fmt.Errorf("prettifying abi: %w", err)
						}

						if err := os.WriteFile(filepath.Join(outdir, t+".abi"), pretty.Bytes(), 0600); err != nil {
							return fmt.Errorf("writing abi: %w", err)
						}

						if err := os.WriteFile(filepath.Join(outdir, t+".bin"), []byte(md.Bins[i]), 0600); err != nil {
							return fmt.Errorf("writing bin: %w", err)
						}
					}
				}
				return nil
			},
		); err != nil {
			return err
		}
		return nil
	},
}
