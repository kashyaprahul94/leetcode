package medium

import "testing"

func TestAddTwoNumbers (t *testing.T) {
	res := addTwoNumbers(1, 2)

	if res != 3 {
		t.Fatalf("Expected %v but got %v", 3, res)
	}
}