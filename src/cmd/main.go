package main

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/k0kubun/pp"
)

func main() {
	filename := "./src/cmd/main.go"

	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		panic("[ERROR] failed to parse Go file. Can your Go file compile?")
	}

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
							pp.Printf("before op: %s\n", expr.Op.String())
							expr = mutator.PlusToMinus(expr)
							pp.Printf("after op: %s\n", expr.Op.String())
						}
					}

				}
			}

		}
	}

}

func hoge() (int, error) {
	return 1 + 2*3, nil
}
