package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/sirupsen/logrus"
)

var (
	log     = logger.New()
	mutated = []token.Pos{}
)

func main() {
	filename := "./src/cmd/main.go"
	log.WithFields(logrus.Fields{
		"filename": filename,
	}).Debug("[main] log start")

	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error_msg": err.Error(),
		}).Error("[main] Failed to parser.ParseFile")
		panic("[ERROR] failed to parse Go file. Can your Go file compile?")
	}
	output := f

	m := mutator.New(log, mutated)

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name != "main" {

				log.WithFields(logrus.Fields{
					"func_name": d.Name.Name,
				}).Debug("[Decl] func is found")

				for _, stmt := range d.Body.List {
					m.MutateStmt(stmt)
				}

			}
		}
	}

	format.Node(os.Stdout, token.NewFileSet(), output)
}

func add(x, y int) int {
	return x + y
}

func hoge(x int) (int, error) {
	y := add(1+2, 3)
	return x + (1 - (2 + y)), nil
}
