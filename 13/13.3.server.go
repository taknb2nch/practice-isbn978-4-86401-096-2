package main

import (
	"fmt"
	"net"
	_ "net/http"
	"net/rpc"
)

type Counter struct {
	count int
}

type Arg struct {
	Increment int
}

type Result struct {
	Value int
}

func (c *Counter) Value(in *Arg, out *Result) error {
	c.count += in.Increment
	out.Value = c.count
	return nil
}

func main() {
	server := rpc.NewServer()
	server.RegisterName("GoCounter", new(Counter))

	l, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println(err)
	}

	server.Accept(l)

	// counter := new(Counter)
	// rpc.Register(counter)
	// rpc.HandleHTTP()

	// l, err := net.Listen("tcp", ":1234")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //go http.Serve(l, nil)
	// http.Serve(l, nil)
}
