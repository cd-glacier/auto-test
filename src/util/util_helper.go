package util

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
