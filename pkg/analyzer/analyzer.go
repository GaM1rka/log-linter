package analyzer

import (
	"go/ast"
	"log"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc: `Checks:
- log message starts with lowercase
- message is English-only
- message contains no emoji/special symbols
- message contains no sensitive data`,
	Run:              run,
	RunDespiteErrors: true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		ce, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		if sel, ok := ce.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok {
				if ident.Name == "slog" {
					switch sel.Sel.Name {
					case "Info":
						log.Print("Uraaaaa")
						pass.Reportf(ce.Pos(), "slog.Info method")
					case "Warn":
						pass.Reportf(ce.Pos(), "slog.Warn method")
					case "Debug":
						pass.Reportf(ce.Pos(), "slog.Debug method")
					case "Error":
						pass.Reportf(ce.Pos(), "slog.Error method")
					}
				}
			}
		}
		return true
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}
	return nil, nil
}
