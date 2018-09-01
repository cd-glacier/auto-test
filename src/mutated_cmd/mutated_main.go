package main

import (
	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/util"
)

var (
	log = logger.New()
)

func main() {
	log.Info("[main] task start")

	// input
	src := "./src/cmd/"
	/*
		operatorType := []operator.Type{
			operator.TO_MAXINT,
		}
	*/

	targetFiles, err := util.FindMutateFile(src)
	if err != nil {
		log.Error("[main] Failed to util.FindMutateFile: %s", err.Error())
	}

	for _, targetFile := range targetFiles {
		targetDir := util.GetDirFromFileName(targetFile)

		util.CreateMutatedDir("mutated_", targetDir)
	}

	log.Info("[main] task finished")
}
