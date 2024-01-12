package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := cli.NewApp()
	app.Name = "ethgen"
	app.Usage = compileCmd.Usage
	app.Description = "This is a library for generating creating strictly typed ethereum modules."
	app.HideVersion = true
	app.Before = func(context *cli.Context) error {
		if context.Bool("verbose") {
			log.SetFlags(0)
		} else {
			log.SetOutput(io.Discard)
		}
		return nil
	}

	app.Action = compileCmd.Action
	app.Commands = []*cli.Command{
		bindCmd,
		compileCmd,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
