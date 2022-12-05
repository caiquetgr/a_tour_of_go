package ch13

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)

	// creating a comparer to ignore the DateAdded field
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
