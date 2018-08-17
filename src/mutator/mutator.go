package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/sirupsen/logrus"
)

type Mutator struct {
	file      *ast.File
	operators []operator.Type
	log       *logrus.Logger
}

func ParseFile(filename string) (*ast.File, error) {
	return parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
}

func New(file *ast.File, operators []operator.Type, log *logrus.Logger) *Mutator {
	return &Mutator{file, operators, log}
}

func (m *Mutator) mutate() (*ast.File, error) {
	m.log.WithFields(logrus.Fields{
		"operator_type": m.operators,
	}).Info("[mutator]")

	return m.file, nil
}
