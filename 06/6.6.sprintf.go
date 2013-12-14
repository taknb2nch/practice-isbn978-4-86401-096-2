package main

import (
	"fmt"
)

func main() {
	f := 1.0
	str := fmt.Sprintf("%T, %T %#v, %d, %v", f, main, main, 42, "hogehoge")
	fmt.Println(str)
}