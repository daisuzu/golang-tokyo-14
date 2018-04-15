package main // OMIT

import (
	"go/ast"

	"honnef.co/go/tools/lint"
)

func (*checker) Check0001(j *lint.Job) {
	fn := func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		j.Errorf(call, "something is wrong")
		return true
	}
	for _, f := range j.Program.Files {
		ast.Inspect(f, fn)
	}
}
