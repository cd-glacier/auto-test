package mutator

import (
	"go/ast"

	"github.com/sirupsen/logrus"
)

func (m *Mutator) MutateDecl(decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		if d.Name.Name != "main" {

			m.log.WithFields(logrus.Fields{
				"func_name": d.Name.Name,
			}).Debug("[Decl] func is found")

			for _, stmt := range d.Body.List {
				m.MutateStmt(stmt)
			}
		}
	}
}
