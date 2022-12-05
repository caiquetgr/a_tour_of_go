package ch13

import "testing"

func Test_addNumbers(t *testing.T) {
	result := AddNumbers(2, 3)
	if result != 5 {
		//t.Error("incorrent result: expected 5, got", result)
		// t.Errorf("incorrent result: expected %d, got %d", 5,result)
    t.Fatal("incorrect result: expected 5, got", result)
	}
}
