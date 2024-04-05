//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/withtally/synceth/codegen"
	"github.com/withtally/synceth/example"
	testexample "github.com/withtally/synceth/testexample/example"
)

func main() {
	tea := "testexample"
	solversion := "0.8.0"
	if err := codegen.GenerateBindings("./artifacts", "./bindings", &codegen.BindingsConfig{
		Fakes:        true,
		FakesVersion: &solversion,
		Handlers: codegen.HandlersConfig{
			Generate: true,
			InputTypes: []codegen.InputType{
				{
					Name: "tx", Type: &example.TestInput{},
				},
				{
					Name: "testtx", Type: &testexample.TestInput{}, Alias: &tea,
				},
			},
		},
		Setup: codegen.SetupConfig{
			InputTypes: []codegen.InputType{
				{
					Name: "i", Type: &example.TestInput{},
				},
			},
		},
	}); err != nil {
		log.Fatalf("running ethgen codegen: %v", err)
	}
}
