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
	if err := codegen.GenerateBindings("./artifacts", "./bindings", &codegen.BindingsConfig{
		Fakes: true,
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
	}); err != nil {
		log.Fatalf("running ethgen codegen: %v", err)
	}
}
