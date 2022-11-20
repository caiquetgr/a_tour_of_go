package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	result, err := timeLimit()
	fmt.Println(result, err)
}

func timeLimit() (int, error) {
	var result int
	var err error

	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	time.Sleep(3 * time.Second)
	return 1, nil
}
