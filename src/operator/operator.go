package operator

import (
	"math"
	"strconv"

	"go/ast"
	"go/token"
)

type Type int

type Operator struct {
	Type    Type
	Literal string
}

const (
	_ = iota
	TO_MAXINT
	ASSIGN_SCOPE
)

var types = []string{
	"ILLEGAL",
	"TO_MAXINT",
	"ASSIGN_SCOPE",
}

func (t Type) String() string {
	s := ""
	if 0 <= t && t < Type(len(types)) {
		s = types[t]
	}
	if s == "" {
		s = "type(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}

func ToMaxInt(node *ast.AssignStmt) {
	bl, ok := node.Rhs[0].(*ast.BasicLit)
	if ok && bl.Kind == token.INT {
		node.Rhs = []ast.Expr{
			&ast.BasicLit{
				Kind:  token.INT,
				Value: strconv.Itoa(math.MaxInt64),
			},
		}
	}
}
