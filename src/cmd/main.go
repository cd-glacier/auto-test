package main

import (
	"go/token"

	"github.com/g-hyoga/auto-test/src/logger"
)

var (
	log     = logger.New()
	mutated = []token.Pos{}
)

func main() {
	log.Info("[main] task start")
	log.Info("[main] task finished")
}
