package main

import (
	"os"

	"honnef.co/go/tools/lint/lintutil"
)

func main() {
	fs := lintutil.FlagSet("mylint")
	fs.Parse(os.Args[1:])

	confs := []lintutil.CheckerConfig{
		{Checker: NewChecker(), ExitNonZero: true},
	}

	lintutil.ProcessFlagSet(confs, fs)
}
