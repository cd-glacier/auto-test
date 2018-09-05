package operator

import (
	"strconv"

	"github.com/rogpeppe/godef/go/ast"
)

type Type int

type Operator struct {
	Type    Type
	Literal string
}

const (
	_ = iota
	TO_MAXINT
)

var types = []string{
	"ILLEGAL",
	"TO_MAXINT",
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

func ToMaxInt(decl ast.FuncDecl) *ast.FuncDecl {
	/*
		ast.Inspect(decl, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				pp.Println(node)
			}

			return true
		})
	*/

	return &decl
}
