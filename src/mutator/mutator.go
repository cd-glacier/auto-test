package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/sirupsen/logrus"
)

type Mutator struct {
	log     *logrus.Logger
	mutated []token.Pos
}

func New(log *logrus.Logger, mutated []token.Pos) *Mutator {
	return &Mutator{log, mutated}
}

func (m *Mutator) ParseFile(filename string) (*ast.File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		m.log.WithFields(logrus.Fields{
			"error_msg": err.Error(),
		}).Error("[main] Failed to parser.ParseFile")
		panic("[ERROR] failed to parse Go file. Can your Go file compile?")
	}
	return f, err
}

func contain(target token.Pos, positions []token.Pos) bool {
	for _, pos := range positions {
		if pos == target {
			return true
		}
	}
	return false
}
