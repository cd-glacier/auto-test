package main

import (
	"go/token"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/sirupsen/logrus"
)

var (
	log     = logger.New()
	mutated = []token.Pos{}
)

func main() {
	src := "./src/cmd"

	copiedDir := prepare(src)
	err := mutate(copiedDir)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error_msg": err.Error(),
		}).Error("[main] failed to mutate")
	}

	postProcess(copiedDir)

	log.Info("[DONE] created mutated go code")
}

func prepare(src string) string {
	copiedDir, err := util.CreateMutatedDir(src)
	if err != nil {
		log.WithFields(logrus.Fields{
			"src":       src,
			"error_msg": err,
		}).Error("[ERROR] failed to create mutated dirs")
		panic("[ERROR] failed to create mutated dirs")
	}

	log.WithFields(logrus.Fields{
		"created_dir": copiedDir,
	}).Debug("[main] create dir for mutation testing")

	return copiedDir
}

func mutate(dir string) error {
	log.WithFields(logrus.Fields{
		"target_dir": dir,
	}).Info("[main] mutate go code")

	filename := "./src/mutated_cmd/mutated_main.go"

	m := mutator.New(log, mutated)
	f, err := m.ParseFile(filename)
	if err != nil {
		return err
	}

	for _, decl := range f.Decls {
		m.MutateDecl(decl)
	}

	return nil
}

func add(x, y int) int {
	return x + y
}

func postProcess(tmp string) {
	err := util.DeleteMuatedDir(tmp)
	if err != nil {
		log.WithFields(logrus.Fields{
			"copiedDir": tmp,
			"error_msg": err,
		}).Error("[ERROR] failed to delete mutated dirs")
	}
}

func hoge(x int) (int, error) {
	y := add(1+2, 3)
	return x + (1 - (2 + y)), nil
}
