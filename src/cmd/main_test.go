package main

import (
	"testing"
)

func TestHoge(t *testing.T) {
	if num, _ := hoge(0); num != -7 {
		t.Fatal("miss")
	}

}
