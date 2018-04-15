package main // OMIT

import (
	"honnef.co/go/tools/lint"
)

func NewChecker() lint.Checker {
	return &checker{}
}

type checker struct{}

func (*checker) Name() string            { return "mylint" }
func (*checker) Prefix() string          { return "ML" }
func (*checker) Init(prog *lint.Program) {}

func (c *checker) Funcs() map[string]lint.Func {
	return map[string]lint.Func{
		"ML0001": c.Check0001,
	}
}
