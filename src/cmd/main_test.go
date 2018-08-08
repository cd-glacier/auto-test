package main

import (
	"testing"
)

func TestMutation(t *testing.T) {
	if num := mutationTest(0); num != 6 {
		t.Fatal("miss")
	}

}
