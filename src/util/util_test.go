package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	e, err := NewEnv([]string{"tmp.g0"}, []string{"hoge"}, "tmp-dir")
	if err != nil {
		t.Fatalf("Failed to NewEnv %s", err.Error())
	}

	tests := []struct {
		inputFilename    string
		inputContent     string
		expectedFilename string
		expectedContent  string
	}{
		{
			e.filenames[0],
			e.contents[0],
			"copied_file",
			e.contents[0],
		},
	}

	// preparation
	err = e.createTestEnv()
	if err != nil {
		t.Fatalf("Failed to createTestEnv %s", err.Error())
	}
	defer e.closeTestEnv()

	for _, test := range tests {
		// execution
		fp := filepath.Join(e.dirname, test.inputFilename)
		err = copyFile(fp, test.expectedFilename)
		if err != nil {
			t.Fatalf("Failed to util.copyFile %s", err.Error())
		}

		defer func(filename string) {
			err = os.Remove(filename)
			if err != nil {
				t.Fatalf("failed to os.Remove %s: %s", filename, err.Error())
			}
		}(test.expectedFilename)

		// inspection
		info, err := os.Stat(test.expectedFilename)
		if err != nil {
			t.Fatalf("failed to os.Stat %s: %s", test.expectedFilename, err.Error())
		}
		if info.Name() != test.expectedFilename {
			t.Fatalf("failed to util.copyFile. copyFile name is invalid %s: %s", test.expectedFilename, err.Error())
		}

		file, err := os.Open(fp)
		defer file.Close()
		buf := make([]byte, len(test.inputContent))
		n, err := file.Read(buf)
		if n == 0 {
			t.Fatalf("failed to util.copyFile. %s content is null", fp)
		}
		if err != nil {
			t.Fatalf("failed to util.copyFile. os.Open %s: %s", fp, err.Error())
		}
		if string(buf) != test.inputContent {
			t.Fatalf("failed to util.copyFile. %s content is invalid. expected: %s, actual: %s.", fp, test.inputContent, string(buf))
		}
	}

}
