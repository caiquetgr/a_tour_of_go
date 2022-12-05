package ch13_test

import (
	"learning_go_book/ch13"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := ch13.AddNumbers(2, 3)
	if result != 5 {
		t.Fatal("incorrect result: expected 5, got", result)
	}
}
