package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
)

func extractMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	return extractString(call.Args[0])
}

func extractString(expr ast.Expr) (string, bool) {
	switch v := expr.(type) {

	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return "", false
		}
		s, err := strconv.Unquote(v.Value)
		if err != nil {
			return "", false
		}
		return s, true

	case *ast.BinaryExpr:
		if v.Op != token.ADD {
			return "", false
		}

		left, ok1 := extractString(v.X)
		right, ok2 := extractString(v.Y)
		if !ok1 || !ok2 {
			return "", false
		}

		return left + right, true
	}

	return "", false
}
