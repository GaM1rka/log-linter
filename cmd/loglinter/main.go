package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"log-linter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
