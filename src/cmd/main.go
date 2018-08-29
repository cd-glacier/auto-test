package main

import (
	"go/format"
	"go/token"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.New()
)

func main() {
	log.Info("[main] task start")

	src := "./src/cmd/"
	operatorType := []operator.Type{
		operator.TO_MAXINT,
	}

	// prepare
	copiedDir, err := util.CreateMutatedDir(src)
	if err != nil {
		log.WithFields(logrus.Fields{
			"src":       src,
			"error_msg": err,
		}).Error("[main] failed to create mutated dirs")
		panic("[main] failed to create mutated dirs")
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

	// mutate
	for _, filename := range filenames {
		f, err := mutator.ParseFile(filename)
		if err != nil {
			panic(err)
		}
		m := mutator.New(f, operatorType, log)
		mutatedFiles := m.Mutate()

		for _, mutatedFile := range mutatedFiles {
			file, err := util.ReWrite(filename)
			if err != nil {
				log.WithFields(logrus.Fields{
					"error_msg": err.Error(),
				}).Error("[main] failed to util.Rewrite")
			}

			format.Node(file, token.NewFileSet(), mutatedFile)
			if err != nil {
				log.WithFields(logrus.Fields{
					"error_msg": err.Error(),
				}).Error("[main] failed to format.Node")
			}

			file.Close()
		}
	}

	log.Info("[main] task finished")
}
