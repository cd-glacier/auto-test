package util

import (
	"errors"
	"os"
	"path/filepath"
)

type TestEnv struct {
	filenames []string
	contents  []string
	dirname   string
}

func NewEnv(filenames, contents []string, dirname string) (*TestEnv, error) {
	if len(filenames) != len(contents) {
		return nil, errors.New("Failed to util.NewEnv. filenames and contents length is invalid")
	}
	e := &TestEnv{}
	e.filenames = filenames
	e.contents = contents
	e.dirname = dirname
	return e, nil
}

func (e *TestEnv) createTestEnv() error {
	if _, err := os.Stat(e.dirname); os.IsNotExist(err) {
		err := os.Mkdir(e.dirname, 0755)
		if err != nil {
			return err
		}
	}

	for i, filename := range e.filenames {
		fp := filepath.Join(e.dirname, filename)
		file, err := os.Create(fp)
		if err != nil {
			return err
		}

		file.Write([]byte(e.contents[i]))

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
