package main

import (
	"fmt"
	"unsafe"
)

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func main() {
	s := "sun"
	fmt.Println(*(*stringStruct)(unsafe.Pointer(&s)))
	for _, b := range s {
		fmt.Println(b)
	}
}
