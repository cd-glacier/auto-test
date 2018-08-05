package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/k0kubun/pp"
)

func main() {
	filename := "./src/cmd/main.go"

	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Errorf("ParseFile(%q)", err)
	}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			pp.Printf("%s func is found!!!\n", d.Name.Name)
		}
	}

	src := "a + 2"
	x, err := parser.ParseExpr(src)
	if err != nil {
		fmt.Errorf("ParseExpr(%q): %v", src, err)
	}

	hoge, ok := x.(*ast.BinaryExpr)
	// sanity check
	if !ok {
		fmt.Errorf("ParseExpr(%q): got %T, want *ast.BinaryExpr", src, x)
	}

	hoge.Op = token.SUB

	/*
		pp.Println(hoge)
		pp.Println(hoge.Op.String())
	*/
}

func hoge() {

}
