package main

import (
	"fmt"
	"time"
)

func main() {
	// a time.Duration, representing 2 hours and 30 minutes:
	d := 2 * time.Hour + 30 * time.Minute
	fmt.Println(d)
}
