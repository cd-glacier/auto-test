package util

import "strings"

func removeBlank(li []string) []string {
	list := []string{}
	for _, l := range li {
		if l != "" {
			list = append(list, l)
		}
	}
	return list
}

func changeLastDirName(path []string, prefix string) []string {
	base := path[:len(path)-1]
	dir := path[len(path)-1]
	base = append(base, prefix+dir)
	return base
}

func isMutationTarget(filename string) bool {
	testNotInclude := !strings.Contains(filename, "test")
	isGoFile := filename[len(filename)-3:] == ".go"

	return testNotInclude && isGoFile
}

func joinPath(path ...string) string {
	return strings.Replace(strings.Join(path, "/"), "//", "/", -1)
}
