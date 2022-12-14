package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")

	var stringSliceType = reflect.TypeOf([]string(nil))

	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)

	fmt.Println(ss) // prints [hello]

}
