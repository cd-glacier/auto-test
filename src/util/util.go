package util

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/g-hyoga/auto-test/src/logger"
	"github.com/sirupsen/logrus"
)

var log = logger.New()

func ReWrite(filename string) (*os.File, error) {
	err := os.Remove(filename)
	if err != nil {
		return nil, err
	}

	return os.Create(filename)
}

func FindMutateFile(src string) ([]string, error) {
	foundFiles := []string{}

	directory, err := os.Open(src)
	if err != nil {
		return foundFiles, err
	}

	objects, err := directory.Readdir(-1)
	if err != nil {
		return foundFiles, err
	}

	for _, obj := range objects {
		if !obj.IsDir() &&
			!strings.Contains(obj.Name(), "_test.go") &&
			strings.Contains(obj.Name(), ".go") {
			foundFiles = append(foundFiles, filepath.Join(src, obj.Name()))
		}
	}

	log.WithFields(logrus.Fields{
		"files": foundFiles,
	}).Debug("[util] found mutate files")

	return foundFiles, nil
}

func DeleteMuatedDir(dir string) error {
	return os.RemoveAll(dir)
}

func CreateMutatedDir(src string) (string, error) {
	prefix := "mutated_"
	base, srcDir := filepath.Split(src)
	destDir := filepath.Join(base, prefix+srcDir)
	return destDir, copyDir(src, destDir, prefix)
}

func copyFile(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	_, err = io.Copy(destfile, file)
	if err == nil {
		srcInfo, err := os.Stat(src)
		if err != nil {
			err = os.Chmod(dest, srcInfo.Mode())
		}
	}
	return nil
}

func copyDir(src, dest, prefix string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(src)
	if err != nil {
		return err
	}

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFileName := filepath.Join(src, obj.Name())
		destFileName := filepath.Join(dest, prefix+obj.Name())

		if obj.IsDir() {
			err = copyDir(srcFileName, destFileName, prefix)
			if err != nil {
				log.WithFields(logrus.Fields{
					"src":       srcFileName,
					"dest":      destFileName,
					"error_msg": err.Error(),
				}).Error("[util] failed to copy dir")
			}
		} else {
			err = copyFile(srcFileName, destFileName)
			if err != nil {
				log.WithFields(logrus.Fields{
					"src":       srcFileName,
					"dest":      destFileName,
					"error_msg": err.Error(),
				}).Error("[util] failed to copy file")
			}
		}
	}
	return nil
}
