package main

import (
	"fmt"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/k0kubun/pp"
)

var (
	log = logger.New()

	src = "./src/cmd/"
)

func main() {
	log.Info("[main] task start")

	// input
	/*
		operatorType := []operator.Type{
			operator.TO_MAXINT,
		}
	*/

	targetPackages, err := util.FindPackages(src)
	if err != nil {
		log.Errorf("[main] Failed to util.FindMutatePackages: %s", err.Error())
		panic(fmt.Sprintf("[main] Failed to util.FindMutatePackages: %s", err.Error()))
	}

	pp.Println(targetPackages)

	/*
		for i, targetFile := range targetFiles {
			targetDir := util.GetDirFromFileName(targetFile)

			prefix := "mutated_" + strconv.Itoa(i) + "_"
			mutateDir, err := util.CreateMutatedDir(prefix, targetDir)
			if err != nil {
				log.Errorf("[main] Failed to util.CreateMutatedDir: %s", err.Error())
				panic(fmt.Sprintf("[main] Failed to util.CreateMutatedDir: %s", err.Error()))
			}
			log.Infof("[main] created '%s' directory", mutateDir)

			log.Infof("[main] start to mutate")
			f, err := mutator.ParseFile(mutateDir)
			// m := mutator.New()

			os.RemoveAll(mutateDir)
		}
	*/

	log.Info("[main] task finished")
}
