package tomaxint

import (
	"testing"
)

func TestProductTwo(t *testing.T) {
	actual := productTwo(3)
	if actual != 6 {
		t.Fatalf("Failed productTwo. actual=%d, expected=%d\n", actual, 6)
	}
}

func TestCaseIsBad(t *testing.T) {
	minusTwenty := productTwo(-10)
	four := productTwo(2)
	if four+minusTwenty < 0 {
		t.Fatalf("hoge")
	}
}
