package util

import (
	"errors"
	"fmt"
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

func FindPackages(src string) ([]string, error) {
	foundPackages := []string{}
	base, _ := filepath.Split(src)

	directory, err := os.Open(src)
	if err != nil {
		return foundPackages, err
	}

	objects, err := directory.Readdir(-1)
	if err != nil {
		log.Errorf("[util] Do you assign file? If you assign file, assign directory.")
		return foundPackages, err
	}

	for _, obj := range objects {
		if obj.IsDir() {
			if !strings.Contains(obj.Name(), "vendor") {
				ps, err := FindPackages(filepath.Join(base, obj.Name()))
				if err != nil {
					return foundPackages, err
				}
				foundPackages = append(foundPackages, ps...)
			}
		} else {
			if !strings.Contains(obj.Name(), "test") &&
				strings.Contains(obj.Name(), ".go") {
				foundPackages = append(foundPackages, src)
				break
			}
		}
	}

	log.WithFields(logrus.Fields{
		"packages": foundPackages,
	}).Debug("[util] found mutate package")

	return foundPackages, nil
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
			!strings.Contains(obj.Name(), "test") &&
			strings.Contains(obj.Name(), ".go") {
			foundFiles = append(foundFiles, filepath.Join(src, obj.Name()))
		}
	}

	log.WithFields(logrus.Fields{
		"files": foundFiles,
	}).Debug("[util] found mutate files")

	return foundFiles, nil
}

func GetDirFromFileName(filename string) (string, error) {
	srcInfo, err := os.Stat(filename)
	if err != nil {
		return "", err
	}

	_, file := filepath.Split(filename)
	if srcInfo.IsDir() && len(file) > 0 {
		filename = filename + "/"
	}

	dir, _ := filepath.Split(filename)

	return dir, nil
}

func DeleteMuatedDir(dir string) error {
	return os.RemoveAll(dir)
}

func CreateMutatedDir(prefix, src string) (string, error) {
	li := removeBlank(strings.Split(src, "/"))
	changedPath := changeLastDirName(li, prefix)

	destDir := strings.Join(changedPath, "/")

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

		if srcFileName == destFileName {
			log.Errorf("[util] failed to copy dir. same name already exist: %s\n", srcFileName)
			return errors.New(fmt.Sprintf("[util] failed to copy dir. same name already exist: %s\n", srcFileName))
		}

		if obj.IsDir() {
			err = copyDir(srcFileName, destFileName, prefix)
			if err != nil {
				log.WithFields(logrus.Fields{
					"src":       srcFileName,
					"dest":      destFileName,
					"error_msg": err.Error(),
				}).Error("[util] failed to copy dir")
				return err
			}
		} else {
			err = copyFile(srcFileName, destFileName)
			if err != nil {
				log.WithFields(logrus.Fields{
					"src":       srcFileName,
					"dest":      destFileName,
					"error_msg": err.Error(),
				}).Error("[util] failed to copy file")
				return err
			}
		}
	}
	return nil
}
