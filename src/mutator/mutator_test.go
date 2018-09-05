package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
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

func parse(t *testing.T, input string) *ast.File {
	f, err := parser.ParseFile(token.NewFileSet(), "main.go", input, parser.AllErrors)
	if err != nil {
		t.Fatalf("Failed to parser.ParseFile. %s\n", err)
	}
	return f
}

func getAValue(t *testing.T, file *ast.File) int {
	v, err := strconv.Atoi(file.Decls[0].(*ast.FuncDecl).Body.List[0].(*ast.AssignStmt).Rhs[0].(*ast.BasicLit).Value)
	if err != nil {
		t.Fatalf("Error getAValue %s", err.Error())
	}
	return v
}

func getBValue(t *testing.T, file *ast.File) int {
	v, err := strconv.Atoi(file.Decls[0].(*ast.FuncDecl).Body.List[1].(*ast.AssignStmt).Rhs[0].(*ast.BasicLit).Value)
	if err != nil {
		t.Fatalf("Error getAValue %s", err.Error())
	}
	return v
}

func TestMutate(t *testing.T) {
	input := parse(t, `
package main

func add() int {
	a := 1
	b := 2
	return a + b
}
`)

	output := []ast.File{
		*parse(t, `
package main

func add() int {
	a := 9223372036854775807
	b := 2 
	return a + b
}
`),
		*parse(t, `
package main

func add() int {
	a := 1 
	b := 9223372036854775807
	return a + b
}
`),
	}

	tests := []struct {
		file         *ast.File
		operators    []operator.Type
		mutatedFiles []ast.File
	}{
		{
			input,
			[]operator.Type{
				operator.TO_MAXINT,
			},
			output,
		},
	}

	for testIndex, tt := range tests {
		m := New(tt.file, tt.operators, logger.New())
		mutatedFiles := m.Mutate()

		if len(mutatedFiles) != len(tt.mutatedFiles) {
			t.Fatalf("invalid mutated files length. expected=%d, actual=%d.\n", len(tt.mutatedFiles), len(mutatedFiles))
		}

		for i, mutatedFile := range mutatedFiles {
			mutatedA := getAValue(t, &mutatedFile)
			expectedA := getAValue(t, &tt.mutatedFiles[i])
			if mutatedA != expectedA {
				t.Fatalf("%dth test invalid A value. actual=%d, expcted=%d.\n", testIndex, mutatedA, expectedA)
			}

			mutatedB := getBValue(t, &mutatedFile)
			expectedB := getBValue(t, &tt.mutatedFiles[i])
			if mutatedB != expectedB {
				t.Fatalf("%dth test invalid B value. actual=%d, expcted=%d.\n", testIndex, mutatedB, expectedB)
			}
		}
	}
}
