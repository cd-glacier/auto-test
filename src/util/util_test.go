package util

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/k0kubun/pp"
)

func TestFindMutateFile(t *testing.T) {
	tests := []struct {
		dir        string
		foundFiles []string
	}{
		{
			"../../testdata/tomaxint/",
			[]string{"../../testdata/tomaxint/tomaxint.go"},
		},
		{
			"./",
			[]string{"./util.go", "./util_helper.go"},
		},
	}

	for _, tt := range tests {
		sort.Slice(tt.foundFiles, func(i, j int) bool { return tt.foundFiles[i] > tt.foundFiles[j] })
		pp.Println(tt.foundFiles)
		foundFiles, err := FindMutateFile(tt.dir)
		if err != nil {
			t.Fatalf("Failed to util.FindMutateFile: %s", err.Error())
		}

		if !reflect.DeepEqual(foundFiles, tt.foundFiles) {
			t.Fatalf("Failed to util.FindMutateFile: actual=%s, expected=%s", foundFiles, tt.foundFiles)
		}
	}
}

func TestGetDirFromFileName(t *testing.T) {
	tests := []struct {
		filename string
		dirname  string
	}{
		{
			"../../testdata/tomaxint/tomaxint.go",
			"../../testdata/tomaxint/",
		},
		{
			"../../testdata/tomaxint",
			"../../testdata/tomaxint/",
		},
		{
			"./../util/util.go",
			"./../util/",
		},
	}

	for _, tt := range tests {
		dirname, err := GetDirFromFileName(tt.filename)
		if err != nil {
			t.Fatalf("Failed to util.GetDirFromFileName: %s\n", err.Error())
		}

		if dirname != tt.dirname {
			t.Fatalf("Failed to util.GetDirFromFileName. actual=%s, expected=%s\n", dirname, tt.dirname)
		}
	}

}

func TestRemoveBlank(t *testing.T) {
	listIncludeBlank := []string{
		"list",
		"",
		"include",
		"blank",
		"",
		"",
	}

	excludedList := removeBlank(listIncludeBlank)
	for _, li := range excludedList {
		if li == "" {
			t.Fatalf("found blank")
		}
	}
}

func TestChangeLastDirName(t *testing.T) {
	path := []string{
		"first",
		"second",
		"last",
	}
	changed := []string{
		"first",
		"second",
		"changed_last",
	}
	actual := changeLastDirName(path, "changed_")

	if !reflect.DeepEqual(actual, changed) {
		t.Fatalf("Failed to changeLastDirName. actual=%s, expected=%s\n", actual, changed)
	}
}

func TestCreateMutatedDir(t *testing.T) {
	tests := []struct {
		testDirName      string
		testFileName     string
		expectedDirName  string
		expectedFileName string
	}{
		{
			"TestCreateMutatedDir",
			"TestCreateMutatedFile",
			"mutated_TestCreateMutatedDir",
			"mutated_TestCreateMutatedFile",
		},
		{
			"./TestCreateMutatedDir",
			"TestCreateMutatedFile",
			"./mutated_TestCreateMutatedDir",
			"mutated_TestCreateMutatedFile",
		},
	}

	for _, tt := range tests {
		err := os.Mkdir(tt.testDirName, 0777)
		if err != nil {
			t.Fatalf("Error to os.Mkdir in TestCreateMutatedDir: %s\n", err.Error())
		}

		_, err = os.Create(filepath.Join(tt.testDirName, tt.testFileName))
		if err != nil {
			remove(t, tt.testDirName)
			t.Fatalf("Error to os.Create in TestCreateMutatedDir: %s\n", err.Error())
		}

		mutatedDir, err := CreateMutatedDir("mutated_", tt.testDirName)
		if err != nil {
			remove(t, tt.testDirName)
			remove(t, mutatedDir)
			t.Fatalf("Error to TestCreateMutatedDir: %s\n", err.Error())
		}
		if mutatedDir != tt.expectedDirName {
			remove(t, tt.testDirName)
			remove(t, mutatedDir)
			t.Fatalf("Failed to TestCreateMutatedDir: actual=%s, expexted=%s\n", mutatedDir, tt.expectedDirName)
		}
		_, err = os.Stat(filepath.Join(tt.expectedDirName, tt.expectedFileName))
		if err != nil {
			remove(t, tt.testDirName)
			remove(t, mutatedDir)
			t.Fatalf("Error to os.Stat in TestCreateMutatedDir: %s\n", err.Error())
		}
		remove(t, tt.testDirName)
		remove(t, mutatedDir)
	}
}

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
