package mutator

import (
	"go/ast"
	"go/token"

	"github.com/sirupsen/logrus"
)

func (m *Mutator) mutateBinaryExpr(expr ast.Expr) *ast.BinaryExpr {
	e, ok := expr.(*ast.BinaryExpr)
	if !ok {
		m.log.WithFields(logrus.Fields{
			"expr": expr,
		}).Error("[Expr] Failed to convert *ast.BinaryExpr")
	}

	switch x := e.X.(type) {
	case *ast.ParenExpr:
		e = m.mutateBinaryExpr(x.X)
	default:
		e.Op = m.exchangePlusMinusToken(e.Op, e.OpPos)
	}

	switch y := e.Y.(type) {
	case *ast.ParenExpr:
		e = m.mutateBinaryExpr(y.X)
	default:
		e.Op = m.exchangePlusMinusToken(e.Op, e.OpPos)
	}

	m.log.WithFields(logrus.Fields{
		"token.Pos": m.mutated,
	}).Debug("[Expr] mutateBinaryExpr")

	return e
}

func (m *Mutator) exchangePlusMinusToken(tok token.Token, pos token.Pos) token.Token {
	if tok == token.ADD && !contain(pos, m.mutated) {
		tok = token.SUB
		m.mutated = append(m.mutated, pos)
	} else if tok == token.SUB && !contain(pos, m.mutated) {
		tok = token.ADD
		m.mutated = append(m.mutated, pos)
	}
	return tok
}
