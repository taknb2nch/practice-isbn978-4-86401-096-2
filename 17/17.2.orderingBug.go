package main

import (
	"fmt"
	//"runtime"
)

func main() {
	// GOMAXPROCSを追加すると期待した動作をしているように思われる
	//runtime.GOMAXPROCS(2)

	var s string
	var flag bool

	go func() {
		fmt.Scanf("%s\n", &s)
		flag = true
	}()

	for !flag {}

	fmt.Printf("%s\n", s)
}

