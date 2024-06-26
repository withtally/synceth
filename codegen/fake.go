package codegen

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const tmplFake = `
// Code generated by github.com/withtally/synceth, DO NOT EDIT.

pragma solidity {{.Version}};

contract Fake{{.Type}} {
	{{range .ABI.Events}}
	event {{.RawName}}({{(event .Inputs)}});
	function fakeEmit{{.RawName}}({{$s := separator ", "}}{{range .Inputs}}{{call $s}}{{(location .Type.String)}} {{.Name}}{{end}}) public {
		emit {{.RawName}}({{$s := separator ", "}}{{range .Inputs}}{{call $s}}{{.Name}}{{end}});
	}
	{{end}}

	{{range .ABI.Methods}}
	{{if not (hasStuct .)}}
	{{$ns := (toCamel .RawName)}}
	{{if gt (len .Outputs) 0}}
		{{(outputVars $ns .Outputs)}}
		function fakeSet{{(toCamel .RawName)}}({{(outputInputs $ns .Outputs)}}) public {
		{{(outputVarAssignments $ns .Outputs)}}
		}
	{{end}}

	function {{.RawName}}({{$s := separator ", "}}{{range .Inputs}}{{call $s}}{{(location .Type.String)}}{{end}}) public view{{if gt (len .Outputs) 0}} returns ({{(outputs .Outputs)}}){{end}} {
		{{if gt (len .Outputs) 0}}return (
			{{$s := separator ", "}}
			{{range $i, $o := .Outputs}}
				{{call $s}}_{{(output $ns (.Type.String) .Name $i)}}
			{{end}}
		);{{end}}
	}
	{{end}}
	{{end}}
}
`

func separator(s string) func() string {
	i := -1
	return func() string {
		i++
		if i == 0 {
			return ""
		}
		return s
	}
}

func location(s string) string {
	if s == "string" || s == "bytes" || strings.HasSuffix(s, "[]") {
		return s + " memory"
	}
	return s
}

func event(args []abi.Argument) string {
	out := ""
	for i, a := range args {
		if len(a.Type.TupleElems) > 0 {
			// TODO: support return structs
		} else {
			if i > 0 {
				out += ", "
			}
			indexed := ""
			if a.Indexed {
				indexed = "indexed "
			}
			out += fmt.Sprintf("%s %s%s", a.Type.String(), indexed, a.Name)
		}
	}
	return out
}

func output(ns, typ, n string, i int) string {
	if n == "" {
		return fmt.Sprintf("ret%s%s%d", strings.Replace(typ, "[]", "Array", 1), ns, i)
	}
	return fmt.Sprintf("%s%s%s", strings.Replace(typ, "[]", "Array", 1), ns, n)
}

func outputInputs(ns string, args []abi.Argument) string {
	out := ""
	for i, a := range args {
		if len(a.Type.TupleElems) > 0 {
			// TODO: support return structs
		} else {
			if i > 0 {
				out += ", "
			}
			out += fmt.Sprintf("%s %s", location(a.Type.String()), output(ns, a.Type.String(), a.Name, i))
		}
	}
	return out
}

func outputVars(ns string, args []abi.Argument) string {
	out := ""
	for i, a := range args {
		if len(a.Type.TupleElems) > 0 {
			// TODO: support return structs
		} else {
			out += fmt.Sprintf("%s private _%s;\n", a.Type.String(), output(ns, a.Type.String(), a.Name, i))
		}
	}
	return out
}

func outputVarAssignments(ns string, args []abi.Argument) string {
	out := ""
	for i, a := range args {
		if len(a.Type.TupleElems) > 0 {
			// TODO: support return structspkg
		} else {
			out += fmt.Sprintf("_%s = %s;\n", output(ns, a.Type.String(), a.Name, i), output(ns, a.Type.String(), a.Name, i))
		}
	}
	return out
}

func outputs(args []abi.Argument) string {
	out := ""
	for i, a := range args {
		if len(a.Type.TupleElems) > 0 {
			// TODO: support return structs
		} else {
			if i > 0 {
				out += ", "
			}
			out += fmt.Sprintf("%s", location(a.Type.String()))
		}
	}
	return out
}

func hasStuct(m abi.Method) bool {
	for _, in := range m.Inputs {
		if len(in.Type.TupleElems) > 0 {
			return true
		}
	}

	for _, out := range m.Outputs {
		if len(out.Type.TupleElems) > 0 {
			return true
		}
	}

	return false
}

type tmplFakeData struct {
	Type    string
	ABI     abi.ABI
	Version string
}

func GenerateFake(typ string, cABI string, pkg string, solversionOverride *string) (string, error) {
	evmABI, err := abi.JSON(strings.NewReader(cABI))
	if err != nil {
		return "", fmt.Errorf("parsing abi: %w", err)
	}

	solversion := "^0.8.3"
	if solversionOverride != nil {
		solversion = *solversionOverride
	}

	data := &tmplFakeData{
		Type:    abi.ToCamelCase(typ),
		ABI:     evmABI,
		Version: solversion,
	}

	buffer := new(bytes.Buffer)
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"event":                event,
		"hasStuct":             hasStuct,
		"outputInputs":         outputInputs,
		"outputs":              outputs,
		"outputVars":           outputVars,
		"outputVarAssignments": outputVarAssignments,
		"location":             location,
		"output":               output,
		"toCamel":              abi.ToCamelCase,
		"separator":            separator,
	}).Parse(tmplFake))
	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}

	return string(buffer.Bytes()), nil
}
