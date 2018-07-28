package main

import "github.com/k0kubun/pp"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func newRange(s, e int) []int {
	slice := []int{}
	for i := s; i < e; i++ {
		slice = append(slice, i)
	}
	return slice
}

func calc(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += fib(num)
	}
	return sum
}

func main() {
	numbers := newRange(0, 10)
	sum := calc(numbers)
	pp.Println(sum)
}
