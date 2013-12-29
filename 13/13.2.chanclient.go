package main

import (
	"code.google.com/p/go.exp/old/netchan"
	"fmt"
)

func main() {
	conn, err := netchan.Import("tcp", "localhost:1234")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	ch := make(chan int)

	err = conn.Import("Counter", ch, netchan.Recv, 1)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Counter: %d\n", <-ch)
	//fmt.Printf("Counter: %d\n", <-ch)
}
