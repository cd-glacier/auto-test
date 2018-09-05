package tomaxint

import "fmt"

func addOne(x int) int {
	y := 1
	return x + y
}

func main() {
	hoge := 1
	foo := 2 + 3
	fmt.Println(hoge + foo)
}
