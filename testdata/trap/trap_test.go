package main

import (
	"testing"
)

func TestReDefinication(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"hogefoo",
			"arg contains hoge",
		},
		{
			"0123foo789",
			"arg is more than 10",
		},
	}

	for _, tt := range tests {
		str := ReDefinication(tt.input)
		if str != tt.output {
			t.Fatalf("Failed to ReDefinication Test: expected=%s, actual=%s", tt.output, str)
		}
	}
}
