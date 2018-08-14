package util

import (
	"os"
	"path/filepath"
)

type TestEnv struct {
	filenames []string
	dirname   string
}

func (e *TestEnv) createTestEnv() error {
	if _, err := os.Stat(e.dirname); os.IsNotExist(err) {
		err := os.Mkdir(e.dirname, 0755)
		if err != nil {
			return err
		}
	}

	for _, filename := range e.filenames {
		fp := filepath.Join(e.dirname, filename)
		file, err := os.Create(fp)
		if err != nil {
			return err
		}

		func() { defer file.Close() }()
	}

	return nil
}

func (e *TestEnv) closeTestEnv() error {
	err := os.RemoveAll(e.dirname)
	if err != nil {
		return err
	}
	return nil
}
