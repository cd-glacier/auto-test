package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"

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

func (m *Mutator) Mutate() []ast.File {
	m.log.WithFields(logrus.Fields{
		"operator_type": m.Operators,
	}).Info("[mutator]")

	var originFile ast.File
	originFile = *m.File

	mutatedFiles := []ast.File{}

	for i, decl := range m.File.Decls {
		file := originFile
		mutatedDecls := m.MutateDecl(decl)
		for _, mutatedDecl := range mutatedDecls {
			file.Decls[i] = mutatedDecl
		}
	}

	if len(mutatedFiles) == 0 {
		m.log.Info("[mutator] mutated file was not found.")
	}

	return mutatedFiles
}

func (m *Mutator) MutateDecl(decl ast.Decl) []ast.Decl {
	decls := []ast.Decl{}

	switch d := decl.(type) {
	case *ast.FuncDecl:
		m.log.WithFields(logrus.Fields{
			"func_name": d.Name.Name,
		}).Debug("[mutator] func is found")

		var copiedDecl ast.FuncDecl
		copiedDecl = *d
		funcDecls := m.MutateFuncDecl(&copiedDecl)
		decls = append(decls, funcDecls...)

	case *ast.GenDecl:
		m.log.WithFields(logrus.Fields{
			"func_name": d.Tok.String(),
		}).Debug("[mutator] GenDecl is found")
	}

	return decls
}

func (m *Mutator) MutateFuncDecl(funcDecl *ast.FuncDecl) []ast.Decl {
	funcDecls := []ast.Decl{}
	copiedDecl := *funcDecl

	ast.Inspect(funcDecl, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			/*
				var copiedNode ast.AssignStmt
				copiedNode = *node
				fmt.Printf("       node address: %p\n", node)
				fmt.Printf("copied node address: %p\n", &copiedNode)
				pp.Printf("node == copiedNode: %s\n", reflect.DeepEqual(node, &copiedNode))
			*/

			operator.ToMaxInt(node)
			pp.Println(funcDecl)
			// pp.Printf("node == copiedNode: %s\n", reflect.DeepEqual(node, &copiedNode))

			funcDecls = append(funcDecls, funcDecl)
			// nodeをもどしてもfuncDeclは変化したまま
			// node = &copiedNode
			// pp.Printf("node == copiedNode: %s\n", reflect.DeepEqual(node, &copiedNode))
			funcDecl = &copiedDecl
		}
		return true
	})
	return funcDecls
}
