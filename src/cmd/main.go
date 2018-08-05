package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"

	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/k0kubun/pp"
)

func main() {
	filename := "./src/cmd/main.go"

	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		panic("[ERROR] failed to parse Go file. Can your Go file compile?")
	}
	output := f

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			pp.Printf("%s func is found!!!\n", d.Name.Name)

			for _, stmt := range d.Body.List {
				switch s := stmt.(type) {
				case *ast.ReturnStmt:
					pp.Println("return stmt!!!")

					for _, expr := range s.Results {
						switch expr := expr.(type) {
						case *ast.BinaryExpr:
							pp.Println("BinaryExpr!!!!")
							pp.Printf("before operator: %s\n", expr.Op.String())
							expr = mutator.PlusToMinus(expr)
							pp.Printf("after operator: %s\n", expr.Op.String())
						}
					}

				}
			}

		}
	}

	format.Node(os.Stdout, token.NewFileSet(), output)
}

func hoge(x int) (int, error) {
	return 2*3 + x, nil
}
