package main

import (
	"go/token"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/sirupsen/logrus"
)

var (
	log     = logger.New()
	mutated = []token.Pos{}
)

func main() {
	src := "./src/cmd"

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

	err = util.DeleteMuatedDir(copiedDir)
	if err != nil {
		log.WithFields(logrus.Fields{
			"copiedDir": copiedDir,
			"error_msg": err,
		}).Error("[ERROR] failed to delete mutated dirs")
	}

	/*
		filename, err := util.FindMutateTarget(copiedDir)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error_msg": err,
			}).Error("[ERROR] not found target go code to mutate")
			panic("[ERROR] not found target go code to mutate")
		}

		log.WithFields(logrus.Fields{
			"filename": filename,
		}).Debug("[main] mutation testing starts")
	*/

	log.Info("[DONE] created mutated go code")
}

func add(x, y int) int {
	return x + y
}

func hoge(x int) (int, error) {
	y := add(1+2, 3)
	return x + (1 - (2 + y)), nil
}
