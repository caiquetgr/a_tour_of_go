package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker3(name string) error {
  var i int
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker3("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	err = ResourceErr{
		Resource: "my resource error",
		Code:     10,
	}

	if err != nil {

		var myErrNotExists ResourceErr
		if errors.As(err, &myErrNotExists) {
			fmt.Println(myErrNotExists)
		}

	}
}

// example Error Is comparison within it
type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}
