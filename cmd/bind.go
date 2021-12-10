package cmd

import (
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/withtally/ethgen/codegen"
)

var bindCmd = &cli.Command{
	Name:  "bind",
	Usage: "create ethereum contract bindings from a abis",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "outdir", Usage: "output directory"},
		&cli.BoolFlag{Name: "fakes", Usage: "generate fakes"},
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

		return codegen.GenerateBindings(path, outdir, &codegen.BindingsConfig{
			Fakes: ctx.Bool("fakes"),
			Handlers: codegen.HandlersConfig{
				Generate: ctx.Bool("handlers"),
			},
		})
	},
}
