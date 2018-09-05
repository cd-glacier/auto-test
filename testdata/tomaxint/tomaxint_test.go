package tomaxint

import "testing"

func TestAddOne(t *testing.T) {
	actual := addOne(3)
	if actual != 4 {
		t.Fatalf("Failed addOne. actual=%d, expected=%d\n", actual, 4)
	}
}
