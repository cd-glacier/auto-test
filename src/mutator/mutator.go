package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"

	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/sirupsen/logrus"
	"github.com/ulule/deepcopier"
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

func (m *Mutator) Mutate() []ast.File {
	m.log.WithFields(logrus.Fields{
		"operator_type": m.Operators,
	}).Info("[mutator]")

	var originFile ast.File
	deepcopier.Copy(m.File).To(&originFile)

	mutatedFiles := []ast.File{}

	for i, decl := range m.File.Decls {
		file := originFile
		mutatedDecls := m.MutateDecl(decl)
		for _, mutatedDecl := range mutatedDecls {
			file.Decls[i] = mutatedDecl
			mutatedFiles = append(mutatedFiles, file)
		}
	}

	return mutatedFiles
}

func (m *Mutator) MutateDecl(decl ast.Decl) []ast.Decl {
	decls := []ast.Decl{}

	switch d := decl.(type) {
	case *ast.FuncDecl:
		m.log.WithFields(logrus.Fields{
			"func_name": d.Name.Name,
		}).Debug("[Decl] func is found")

		var declBeforeMutate ast.FuncDecl
		declBeforeMutate = *d
		ast.Inspect(d, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				// copiedNode := *node
				bl, ok := node.Rhs[0].(*ast.BasicLit)
				if ok && bl.Kind == token.INT {
					node.Rhs = []ast.Expr{
						&ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(math.MaxInt64),
						},
					}
					// ここまではちゃんと機能してる
					// appendができない
					decls = append(decls, d)
					// *node = copiedNode

					*d = declBeforeMutate
				}
			}
			return true
		})

	case *ast.GenDecl:
		m.log.WithFields(logrus.Fields{
			"func_name": d.Tok.String(),
		}).Debug("[Decl] GenDecl is found")
	}
	return decls
}
