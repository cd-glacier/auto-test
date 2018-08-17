package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/k0kubun/pp"
)

func TestMutate(t *testing.T) {
	input := `
package main

func add() int {
	a := 1
	a += 1
	return a
}
`

	f, err := parser.ParseFile(token.NewFileSet(), "main.go", input, parser.AllErrors)
	if err != nil {
		t.Fatalf("Failed to parser.ParseFile. %s\n", err)
	}

	tests := []struct {
		code        *ast.File
		operators   []operator.Type
		mutatedCode *ast.File
	}{
		{
			f,
			[]operator.Type{
				operator.TO_MAXINT,
			},
			f,
		},
	}

	for _, test := range tests {
		m := New(test.code, test.operators, logger.New())
		mutatedFile, err := m.mutate()
		if err != nil {
			t.Fatalf("Failed to mutator.mutate. %s\n", err.Error())
		}

		pp.Println(mutatedFile)

	}
}
