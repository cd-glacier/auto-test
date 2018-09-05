package main

import (
	"fmt"
	"go/format"
	"go/token"
	"os"
	"os/exec"
	"strings"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/g-hyoga/auto-test/src/mutator"
	"github.com/g-hyoga/auto-test/src/operator"
	"github.com/g-hyoga/auto-test/src/util"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.New()

	src = "./testdata/"
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
		targetDir, err := util.GetDirFromFileName(targetPackage)
		if err != nil {
			log.Errorf("[main] Failed to util.GetDirFromFileName: %s", err.Error())
			panic(fmt.Sprintf("[main] Failed to util.GetDirFromFileName: %s", err.Error()))
		}

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
			if len(mutatedFiles) == 0 {
				break
			}

			file, err := util.ReWrite(targetFile)
			if err != nil {
				log.Errorf("[main] Failed to util.ReWrite: %s", err.Error())
			}
			err = format.Node(file, token.NewFileSet(), &mutatedFiles[0])
			if err != nil {
				log.Errorf("[main] Failed to format.Node: %s", err.Error())
			}

			test(targetFile)

			file.Close()
		}

		os.RemoveAll(mutateDir)
	}

	log.Info("[main] task finished")
}

func test(dir string) {
	log.WithFields(logrus.Fields{
		"test_target": dir,
	}).Info("[main] start to run test")

	packageName, err := util.GetDirFromFileName(dir)
	if err != nil {
		log.Errorf("[main] Failed to util.GetDirFromFileName: %s", err.Error())
	}

	cmd := exec.Command("go", "test", "-v", packageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "--- FAIL:") {
			log.WithFields(logrus.Fields{
				"output": "\n" + string(out),
			}).Info("[main] failed your test")
		} else {
			log.WithFields(logrus.Fields{
				"dir":       dir,
				"error_msg": err,
				"output":    "\n" + string(out),
			}).Error("[main] failed to run test")
		}
	}

	log.Info("[main] finish to run test")
}
