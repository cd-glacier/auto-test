package main

import (
	"reflect"
	"testing"
)

func TestNewRange(t *testing.T) {
	if !reflect.DeepEqual(newRange(0, 3), []int{0, 1, 2}) {
		t.Fatalf("Failed to newRange")
	}
}

func TestFib(t *testing.T) {
	if fib(0) != 0 {
		t.Fatalf("Failed to fib")
	}

	if fib(5) != 5 {
		t.Fatalf("Failed to fib")
	}
}

func TestCalc(t *testing.T) {
	if calc([]int{1, 2}) != 2 {
		t.Fatalf("Failed to calc")
	}
}
