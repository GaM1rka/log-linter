package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/GaM1rka/log-linter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
