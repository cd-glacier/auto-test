package main

import (
	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/k0kubun/pp"
)

var (
	log = logger.New()
)

func main() {
	log.Info("[main] task start")

	operatorType := []operator.Type{
		operator.TO_MAXINT,
	}
	f, err := mutator.ParseFile("./src/cmd/main.go")
	if err != nil {
		panic(err)
	}
	m := mutator.New(f, operatorType, log)
	m.Mutate()
	pp.Println(m.File)

	log.Info("[main] task finished")
}
