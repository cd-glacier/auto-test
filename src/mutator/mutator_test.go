package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/operator"
)

func TestParseFile(t *testing.T) {
	_, err := ParseFile("./mutator_test.go")
	if err != nil {
		t.Fatalf("Failed to mutatot.ParseFile. %s\n", err)
	}
}

func TestMutate(t *testing.T) {
	input := `
package main

func add() int {
	a := 1
	return a + 2
}
`

	f, err := parser.ParseFile(token.NewFileSet(), "main.go", input, parser.AllErrors)
	if err != nil {
		t.Fatalf("Failed to parser.ParseFile. %s\n", err)
	}

	tests := []struct {
		file        *ast.File
		operators   []operator.Type
		mutatedFile *ast.File
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
		m := New(test.file, test.operators, logger.New())
		_, err := m.Mutate()
		if err != nil {
			t.Fatalf("Failed to mutator.mutate. %s\n", err.Error())
		}

	}
}
