package analyzer

import (
	"go/ast"

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
					case "Info", "Warn", "Error", "Debug":
						runRules(pass, ce)
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

func runRules(pass *analysis.Pass, ce *ast.CallExpr) {
	msg, ok := extractMessage(ce)
	if !ok {
		return
	}

	if errMsg, ok := checkStartsWithLowercase(msg); !ok {
		pass.Reportf(ce.Pos(), errMsg)
	}
	if errMsg, ok := checkEnglishOnly(msg); !ok {
		pass.Reportf(ce.Pos(), errMsg)
	}
	if errMsg, ok := checkNoEmojiOrSpecial(msg); !ok {
		pass.Reportf(ce.Pos(), errMsg)
	}
	if errMsg, ok := checkNoSensitive(msg); !ok {
		pass.Reportf(ce.Pos(), errMsg)
	}
}
