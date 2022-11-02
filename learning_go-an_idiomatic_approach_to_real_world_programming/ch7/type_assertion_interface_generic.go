package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	//i3 := i.(string) // panic
	//fmt.Println(i3)

	i4, ok := i.(int) // also panics! must really be the exact type MyInt, we can use the "ok" idiom
	if !ok {
		fmt.Println(fmt.Errorf("unexpected type for %v", i).Error())
	} else {
		fmt.Println(i4)
	}

	switch j := i.(type) {
	case nil:
	case int:
	case MyInt:
		fmt.Println(j)
	case string:
	case io.Reader:
	case bool, rune:
	default:
		// becomes of type interface{}
		fmt.Println("ih")
	}
}
