package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // time is a moment of time, with timezone
	fmt.Println(now)

	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		panic("eita")
	}
	fmt.Println(t)
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	fmt.Println(t.Month())
	fmt.Println(t.AddDate(0,0,-2))

}

