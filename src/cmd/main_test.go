package main

import (
	"testing"
)


func TestFib(t *testing.T) {
	if fib(0) != 0 {
		t.Fatalf("Failed to fib")
	}

	if fib(5) != 5 {
		t.Fatalf("Failed to fib")
	}
}
