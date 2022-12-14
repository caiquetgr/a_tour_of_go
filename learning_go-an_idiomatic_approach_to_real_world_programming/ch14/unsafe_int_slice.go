package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := []int{10, 20, 30}
	sHdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // prints 3
	fmt.Println(sHdr.Cap) // prints 3

	intByteSize := unsafe.Sizeof(s[0])
	fmt.Println(intByteSize)

	for i := 0; i < sHdr.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(sHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	runtime.KeepAlive(s)
}
