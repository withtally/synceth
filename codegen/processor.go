package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"reflect"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var ErrNoEvents = errors.New("no events")

const tmplProcessor = `
// Code generated by github.com/withtally/synceth, DO NOT EDIT.

package {{.Package}}

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	{{range $type := .InputTypes}}
		{{generateAlias .Alias}} "{{.PkgPath}}"
	{{end}}
)

{{range $handler := .Processors}}
	type {{.Type}}Processor interface {
		Setup(ctx context.Context, address common.Address, eth interface {
			ethereum.ChainReader
			ethereum.ChainStateReader
			ethereum.TransactionReader
			bind.ContractBackend
		}, {{$s := separator ", "}}{{range $type := $.SetupInputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error
    	Initialize(ctx context.Context, start uint64, {{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error
		{{range .Events}}
			Process{{.Normalized.Name}}(ctx context.Context, e {{$handler.Type}}{{.Normalized.Name}}) (func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error, error)
		{{end}}
		mustEmbedBase{{.Type}}Processor()
	}

	type Base{{.Type}}Processor struct {
		Address  common.Address
		ABI      abi.ABI
		Contract *{{.Type}}
		Eth      interface {
			ethereum.ChainReader
			ethereum.ChainStateReader
			ethereum.TransactionReader
			bind.ContractBackend
		}
	}

	func (h *Base{{.Type}}Processor) Setup(ctx context.Context, address common.Address, eth interface {
		ethereum.ChainReader
		ethereum.ChainStateReader
		ethereum.TransactionReader
		bind.ContractBackend
	}, {{$s := separator ", "}}{{range $type := $.SetupInputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error {
		contract, err := New{{.Type}}(address, eth)
		if err != nil {
			return fmt.Errorf("new {{.Type}}: %w", err)
		}

		abi, err := abi.JSON(strings.NewReader(string({{.Type}}ABI)))
		if err != nil {
			return fmt.Errorf("parsing {{.Type}} abi: %w", err)
		}

		h.Address = address
		h.ABI = abi
		h.Contract = contract
		h.Eth = eth
		return nil
	}

	func (h *Base{{.Type}}Processor) ProcessElement(p interface{}) func(context.Context, types.Log) (func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error, error) {
		return func(ctx context.Context, vLog types.Log) (func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error, error) {
			switch vLog.Topics[0].Hex() {
			{{range .Events}}
			case h.ABI.Events["{{.Normalized.Name}}"].ID.Hex():
				e := {{$handler.Type}}{{.Normalized.Name}}{}
				if err := h.UnpackLog(&e, "{{.Normalized.Name}}", vLog); err != nil {
					return nil, fmt.Errorf("unpacking {{.Normalized.Name}}: %w", err)
				}

				e.Raw = vLog
				cb, err := p.({{$handler.Type}}Processor).Process{{.Normalized.Name}}(ctx, e);
				if err != nil {
					return nil, fmt.Errorf("processing {{.Normalized.Name}}: %w", err)
				}

				return cb, nil
			{{end}}
			}
			return func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error { return nil }, nil
		}
	}

	func (h *Base{{$handler.Type}}Processor) UnpackLog(out interface{}, event string, log types.Log) error {
		if len(log.Data) > 0 {
			if err := h.ABI.UnpackIntoInterface(out, event, log.Data); err != nil {
				return err
			}
		}
		var indexed abi.Arguments
		for _, arg := range h.ABI.Events[event].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}
		return abi.ParseTopics(out, indexed, log.Topics[1:])
	}

	func (h *Base{{$handler.Type}}Processor) Initialize(ctx context.Context, start uint64, {{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error {
		return nil
	}

	{{range .Events}}
		func (h *Base{{$handler.Type}}Processor) Process{{.Normalized.Name}}(ctx context.Context, e {{$handler.Type}}{{.Normalized.Name}}) (func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error, error) {
			return func({{$s := separator ", "}}{{range $type := $.InputTypes}}{{call $s}}{{$type.Name}} {{formatPointer $type.Kind}}{{$type.Ident}}{{end}}) error { return nil }, nil
		}
	{{end}}

	func (h *Base{{$handler.Type}}Processor) mustEmbedBase{{$handler.Type}}Processor() {}
{{end}}
`

type tmplEventData struct {
	Original   abi.Event // Original event as parsed by the abi package
	Normalized abi.Event // Normalized version of the parsed fields
}

type tmplProcessorData struct {
	Type   string
	Events map[string]*tmplEventData
}

type inputType struct {
	Alias   *string
	Name    string
	Ident   string
	Kind    reflect.Kind
	PkgPath string
}

type tmplData struct {
	Package         string
	InputTypes      []inputType
	SetupInputTypes []inputType
	Processors      map[string]*tmplProcessorData
}

func GenerateProcessor(types []string, abis []string, pkg string, inputs []InputType, setupInputs []InputType) (string, error) {
	var handlers = make(map[string]*tmplProcessorData)
	var n int

	var inputTypes []inputType
	for _, v := range inputs {
		t := reflect.TypeOf(v.Type)
		tv := indirect(t)

		i := tv.String()
		if v.Alias != nil {
			s := strings.Split(i, ".")
			i = fmt.Sprintf("%s.%s", *v.Alias, s[1])
		}

		inputTypes = append(inputTypes, inputType{
			Alias:   v.Alias,
			Kind:    t.Kind(),
			Name:    v.Name,
			Ident:   i,
			PkgPath: tv.PkgPath(),
		})
	}

	var setupInputTypes []inputType
	for _, v := range setupInputs {
		t := reflect.TypeOf(v.Type)
		tv := indirect(t)

		i := tv.String()
		if v.Alias != nil {
			s := strings.Split(i, ".")
			i = fmt.Sprintf("%s.%s", *v.Alias, s[1])
		}

		setupInputTypes = append(setupInputTypes, inputType{
			Alias:   v.Alias,
			Kind:    t.Kind(),
			Name:    v.Name,
			Ident:   i,
			PkgPath: tv.PkgPath(),
		})
	}

	for i, typ := range types {
		var events = make(map[string]*tmplEventData)
		evmABI, err := abi.JSON(strings.NewReader(abis[i]))
		if err != nil {
			return "", fmt.Errorf("parsing abi: %w", err)
		}

		for _, original := range evmABI.Events {
			// Skip anonymous events as they don't support explicit filtering
			if original.Anonymous {
				continue
			}
			// Normalize the event for capital cases and non-anonymous outputs
			normalized := original
			normalized.Name = abi.ToCamelCase(original.Name)

			events[original.Name] = &tmplEventData{Original: original, Normalized: normalized}
		}

		if len(events) > 0 {
			handlers[typ] = &tmplProcessorData{
				Type:   abi.ToCamelCase(typ),
				Events: events,
			}
			n++
		}
	}

	if n == 0 {
		return "", ErrNoEvents
	}

	data := &tmplData{
		InputTypes:      inputTypes,
		SetupInputTypes: setupInputTypes,
		Package:         pkg,
		Processors:      handlers,
	}

	buffer := new(bytes.Buffer)
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatPointer": formatPointer,
		"separator":     separator,
		"generateAlias": generateAlias,
	}).Parse(tmplProcessor))
	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}

	handler, err := format.Source(buffer.Bytes())
	if err != nil {
		return "", err
	}

	return string(handler), nil
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func formatPointer(k reflect.Kind) string {
	if k == reflect.Ptr {
		return "*"
	}

	return ""
}

func generateAlias(alias *string) string {
	if alias != nil {
		return fmt.Sprintf("%s ", *alias)
	}

	return ""
}
