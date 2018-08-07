package mutator

import (
	"go/ast"

	"github.com/sirupsen/logrus"
)

func (m *Mutator) MutateStmt(stmt ast.Stmt) ast.Stmt {
	m.log.WithFields(logrus.Fields{
		"stmt": stmt,
	}).Debug("[Stmt] MutateStmt")

	switch s := stmt.(type) {
	case *ast.AssignStmt:
		for _, expr := range s.Rhs {
			expr = m.mutateReturnStmt(expr)
		}
	case *ast.ReturnStmt:
		for _, expr := range s.Results {
			expr = m.mutateReturnStmt(expr)
		}
	}
	return stmt
}

func (m *Mutator) mutateReturnStmt(expr ast.Expr) ast.Expr {
	m.log.WithFields(logrus.Fields{
		"expr": expr,
	}).Debug("[Stmt] mutateReturnStmt")

	switch e := expr.(type) {
	case *ast.BinaryExpr:
		e = m.mutateBinaryExpr(e)
	}

	return expr
}
