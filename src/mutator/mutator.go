package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"

	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
)

type Mutator struct {
	File      *ast.File
	Operators []operator.Type
	log       *logrus.Logger
}

func ParseFile(filename string) (*ast.File, error) {
	return parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
}

func New(file *ast.File, operators []operator.Type, log *logrus.Logger) *Mutator {
	return &Mutator{file, operators, log}
}

func (m *Mutator) Mutate() (*ast.File, error) {
	m.log.WithFields(logrus.Fields{
		"operator_type": m.Operators,
	}).Info("[mutator]")

	for _, decl := range m.File.Decls {
		MutateDecl(decl, m.log)
	}

	return m.File, nil
}

func MutateDecl(decl ast.Decl, log *logrus.Logger) []ast.Decl {
	decls := []ast.Decl{}

	switch d := decl.(type) {
	case *ast.FuncDecl:
		log.WithFields(logrus.Fields{
			"func_name": d.Name.Name,
		}).Debug("[Decl] func is found")

		ast.Inspect(d, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				node.Rhs = []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: strconv.Itoa(math.MaxInt64),
					},
				}
				decls = append(decls, d)
			}
			return true
		})

	case *ast.GenDecl:
		log.WithFields(logrus.Fields{
			"func_name": d.Tok.String(),
		}).Debug("[Decl] GenDecl is found")
	}

	pp.Println(decls)
	return decls
}

/*
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
*/
