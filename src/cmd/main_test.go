package main

import "testing"

func TestHoge(t *testing.T) {
	if num, _ := hoge(0); num != 6 {
		t.Fatal("miss")
	}

	if num, _ := hoge(2); num != 8 {
		t.Fatal("miss")
	}
}
