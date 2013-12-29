package main

import (
	"fmt"
	"net/rpc"
)

type Arg struct {
	Increment int
}

type Result struct {
	Value int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		fmt.Println(err)
		return
	}

	var r Result
	client.Call("GoCounter.Value", &Arg{1}, &r)

	fmt.Printf("%d\n", r.Value)
}
