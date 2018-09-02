package main

import (
	"fmt"
	"go/format"
	"go/token"
	"os"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.New()

	src = "./src/cmd/"
)

func main() {
	log.Info("[main] task start")

	operatorType := []operator.Type{
		operator.TO_MAXINT,
	}

	targetPackages, err := util.FindPackages(src)
	if err != nil {
		log.Errorf("[main] Failed to util.FindMutatePackages: %s", err.Error())
		panic(fmt.Sprintf("[main] Failed to util.FindMutatePackages: %s", err.Error()))
	}

	for _, targetPackage := range targetPackages {
		targetDir := util.GetDirFromFileName(targetPackage)

		prefix := "mutated_"
		mutateDir, err := util.CreateMutatedDir(prefix, targetDir)
		if err != nil {
			log.Errorf("[main] Failed to util.CreateMutatedDir: %s", err.Error())
			panic(fmt.Sprintf("[main] Failed to util.CreateMutatedDir: %s", err.Error()))
		}
		log.Infof("[main] created '%s' directory", mutateDir)

		targetFiles, err := util.FindMutateFile(mutateDir)
		if err != nil {
			log.Errorf("[main] Failed to util.FindMutateFile: %s", err.Error())
			panic(fmt.Sprintf("[main] Failed to util.FindMutateFile: %s", err.Error()))
		}
		pp.Println(targetFiles)

		for _, targetFile := range targetFiles {
			f, err := mutator.ParseFile(targetFile)
			if err != nil {
				log.Errorf("[main] Failed to mutator.ParseFile: %s", err.Error())
				panic(fmt.Sprintf("[main] Failed to mutator.ParseFile: %s", err.Error()))
			}
			m := mutator.New(f, operatorType, log)

			log.WithFields(logrus.Fields{
				"file":          targetFile,
				"operatorTypes": operatorType,
			}).Infof("[main] start to mutate")
			mutatedFiles := m.Mutate()

			file, err := util.ReWrite(targetFile)
			if err != nil {
				log.Errorf("[main] Failed to util.ReWrite: %s", err.Error())
			}
			err = format.Node(file, token.NewFileSet(), &mutatedFiles[0])
			if err != nil {
				log.Errorf("[main] Failed to format.Node: %s", err.Error())
			}

			file.Close()
		}

		os.RemoveAll(mutateDir)
	}

	log.Info("[main] task finished")
}
