package parser

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type solidityListener struct {
	*BaseSolidityListener
	Constraints []string
}

func (s *solidityListener) EnterVersionConstraint(ctx *VersionConstraintContext) {
	s.Constraints = append(s.Constraints, ctx.GetText())
}

type VersionConstraint struct {
	semver.Constraints
}

func NewVersionConstraint(data string) (*VersionConstraint, error) {
	input := antlr.NewInputStream(data)
	lexer := NewSolidityLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewSolidityParser(stream)
	p.BuildParseTrees = true
	tree := p.SourceUnit()
	l := solidityListener{}
	antlr.ParseTreeWalkerDefault.Walk(&l, tree)

	s := strings.Join(l.Constraints, ",")
	c, err := semver.NewConstraint(s)
	if err != nil {
		return nil, err
	}

	return &VersionConstraint{*c}, nil
}
