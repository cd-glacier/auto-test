package mutator

import (
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

func contain(target token.Pos, positions []token.Pos) bool {
	for _, pos := range positions {
		if pos == target {
			return true
		}
	}
	return false
}
