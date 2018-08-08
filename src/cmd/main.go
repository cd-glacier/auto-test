package main

import (
	"go/format"
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

	filenames, mutatedDir := prepare(src)

	for _, filename := range filenames {
		err := mutate(filename)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error_msg": err.Error(),
			}).Error("[main] failed to mutate")
		}
	}

	postProcess(mutatedDir)

	log.Info("[DONE] created mutated go code")
}

func prepare(src string) ([]string, string) {
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

	filenames, err := util.FindMutateFile(copiedDir)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error_msg": err.Error(),
		}).Error("[main] failed to find mutate file")
	}

	return filenames, copiedDir
}

func mutate(filename string) error {
	log.WithFields(logrus.Fields{
		"filename": filename,
	}).Info("[main] mutate go code")

	m := mutator.New(log, mutated)
	f, err := m.ParseFile(filename)
	if err != nil {
		return err
	}

	for _, decl := range f.Decls {
		m.MutateDecl(decl)
	}

	file, err := util.ReWrite(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	format.Node(file, token.NewFileSet(), f)
	if err != nil {
		return err
	}
	return nil
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

func mutationTest(y int) int {
	x := 3 + y
	return 1 + 2 + x
}
