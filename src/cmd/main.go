package main

import (
	"go/format"
	"go/token"
	"os/exec"
	"strings"

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
	mutatedDir = "./" + mutatedDir

	for _, filename := range filenames {
		err := mutate(filename)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error_msg": err.Error(),
			}).Error("[main] failed to mutate")
		}
	}

	test(mutatedDir)

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

func test(dir string) {
	cmd := exec.Command("go", "test", "-v", dir)
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
			}).Error("[ERROR] failed to run test")
		}

	}

	log.Info("[main] finish to run test")
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
