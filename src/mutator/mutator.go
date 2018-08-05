package mutator

import (
	"go/ast"
	"go/token"
)

func PlusToMinus(expr *ast.BinaryExpr) *ast.BinaryExpr {
	if expr.Op == token.ADD {
		expr.Op = token.SUB
	}

	return expr
}
