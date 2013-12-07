package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "A Go variable"
	// http://golang.org/pkg/unsafe/#Pointer
	addr := unsafe.Pointer(&str)
	//addr := &str
	fmt.Printf("The address of str is %d\n", addr)

	str2 := (*string)(addr)
	fmt.Printf("String constructed form pointer: %s\n", *str2)

	address := uintptr(addr)
	address += 4

	//
	str3 := (*string)(unsafe.Pointer(address))
	fmt.Printf("String constructed from pointer: %s\n", *str3)
}
