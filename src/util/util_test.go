package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	// preparation
	e, err := NewEnv([]string{"tmp.g0"}, []string{"hoge"}, "tmp-dir")
	if err != nil {
		t.Fatalf("Failed to NewEnv %s", err.Error())
	}
	err = e.createTestEnv()
	if err != nil {
		t.Fatalf("Failed to createTestEnv %s", err.Error())
	}
	defer e.closeTestEnv()

	// execution
	fp := filepath.Join(e.dirname, e.filenames[0])
	filename := "copied_file"
	err = copyFile(fp, filename)
	if err != nil {
		t.Fatalf("Failed to util.copyFile %s", err.Error())
	}

	defer func(filename string) {
		err = os.Remove(filename)
		if err != nil {
			t.Fatalf("failed to os.Remove %s: %s", filename, err.Error())
		}
	}(filename)

	// inspection
	info, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("failed to os.Stat %s: %s", filename, err.Error())
	}
	if info.Name() != filename {
		t.Fatalf("failed to util.copyFile. copyFile name is invalid %s: %s", filename, err.Error())
	}

	file, err := os.Open(fp)
	defer file.Close()
	buf := make([]byte, len(e.contents[0]))
	n, err := file.Read(buf)
	if n == 0 {
		t.Fatalf("failed to util.copyFile. %s content is null", fp)
	}
	if err != nil {
		t.Fatalf("failed to util.copyFile. os.Open %s: %s", fp, err.Error())
	}
	if string(buf) != e.contents[0] {
		t.Fatalf("failed to util.copyFile. %s content is invalid. expected: %s, actual: %s.", fp, e.contents[0], string(buf))
	}
}
