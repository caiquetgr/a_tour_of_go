package main

import (
	"fmt"
	"reflect"
)

type Foo struct{}

func main() {
	v := 1
	vType := reflect.TypeOf(v)

  fmt.Println("vType name:",vType.Name())
  fmt.Println("vType kind:",vType.Kind())

	f := Foo{}
	ft := reflect.TypeOf(f)

  fmt.Println("ft name: ",ft.Name())
	fmt.Println("ft kind:", ft.Kind())

	pv := reflect.TypeOf(&v)
	fmt.Println("pv name:", pv.Name()) // some types like slice or pointers don't have names
	fmt.Println("pv kind:", pv.Kind())

}
