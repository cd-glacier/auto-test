package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	e := TestEnv{filenames: []string{"tmp.g0", "tmp_test.g0"}, dirname: "tmp-dir"}
	err := e.createTestEnv()
	if err != nil {
		t.Fatalf("Failed to createTestEnv %s", err.Error())
	}
	defer e.closeTestEnv()

	fp := filepath.Join(e.dirname, e.filenames[0])
	filename := "copied_file"
	err = copyFile(fp, filename)
	if err != nil {
		t.Fatalf("Failed to copyFile %s", err.Error())
	}

	defer func(filename string) {
		err = os.Remove(filename)
		if err != nil {
			t.Fatalf("failed to os.Remove %s: %s", filename, err.Error())
		}
	}(filename)

	info, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("failed to os.Stat %s: %s", filename, err.Error())
	}
	if info.Name() != filename {
		t.Fatalf("failed to copyFile. copyFile name is invalid %s: %s", filename, err.Error())
	}

}
