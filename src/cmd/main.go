package main

import (
	"fmt"
	"go/ast"
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
		log.WithFields(logrus.Fields{
			"directory":     mutateDir,
			"operatorTypes": operatorType,
		}).Infof("[main] created '%s' directory", mutateDir)

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
			log.Infof("[main] created %d mutated files", len(mutatedFiles))

			passed, err := mutationTest(targetFile, mutatedFiles)
			if err != nil {
				log.Errorf("[main] Failed to mutationTest: %s", err.Error())
			}

			log.Infof("[main] %d/%d passed tests are found.", len(passed), len(mutatedFiles))

		}

		os.RemoveAll(mutateDir)
	}

	log.Info("[main] task finished")
}

func mutationTest(fileBeforeMutation string, mutatedFiles []ast.File) ([]ast.File, error) {
	passedTests := []ast.File{}

	for _, mutatedFile := range mutatedFiles {
		file, err := util.ReWrite(fileBeforeMutation)
		if err != nil {
			log.Errorf("[main] Failed to util.ReWrite: %s", err.Error())
			return passedTests, err
		}
		err = format.Node(file, token.NewFileSet(), &mutatedFile)
		if err != nil {
			log.Errorf("[main] Failed to format.Node: %s", err.Error())
			return passedTests, err
		}

		if passed := test(fileBeforeMutation); passed {
			passedTests = append(passedTests, mutatedFile)
		}

		file.Close()
	}
	return passedTests, nil
}

func test(file string) bool {
	log.WithFields(logrus.Fields{
		"test_target": file,
	}).Info("[main] start to run test")

	packageName, err := util.GetDirFromFileName(file)
	if err != nil {
		log.Errorf("[main] Failed to util.GetDirFromFileName: %s", err.Error())
	}

	cmd := exec.Command("go", "test", "-v", packageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "--- PASS:") {
			// TODO: count test case
			return true
		} else if strings.Contains(string(out), "--- FAIL:") {
			log.WithFields(logrus.Fields{
				"output": "\n" + string(out),
			}).Debug("[main] failed your test")
			return false
		} else {
			log.WithFields(logrus.Fields{
				"package":   packageName,
				"error_msg": err,
				"output":    "\n" + string(out),
			}).Error("[main] failed to run test")
		}
	}

	log.Info("[main] finish to run test")

	if strings.Contains(string(out), "--- PASS:") {
		return true
	}

	return false
}
