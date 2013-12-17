package main

import (
	"fmt"
	"runtime"
	"time"
)

type RequestType int

const (
	Set = iota
	Get
	BeginTransaction
	EndTransaction
)

type Request struct {
	requestType RequestType
	key         int
	value       string
	ret         chan string
	transaction chan Request
}

func get(m chan Request, key int) string {
	result := make(chan string)
	m <- Request{Get, key, "", result, nil}
	return <-result
}
func set(m chan Request, key int, value string) {
	m <- Request{Set, key, value, nil, nil}
}
func beginTransaction(m chan Request) chan Request {
	t := make(chan Request)
	m <- Request{BeginTransaction, 0, "", nil, t}
	return t
}
func endTransaction(m chan Request) {
	m <- Request{EndTransaction, 0, "", nil, nil}
}

func HandleRequests(m map[int]string, c chan Request) {
	for {
		req := <-c

		switch req.requestType {
		case Get:
			req.ret <- m[req.key]
		case Set:
			m[req.key] = req.value
		case BeginTransaction:
			HandleRequests(m, req.transaction)
		case EndTransaction:
			return
		}
	}
}

func runMap(c chan Request) {
	m := make(map[int]string)
	HandleRequests(m, c)
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(5))

	m := make(chan Request)

	go runMap(m)

	go go1(m)
	go go2(m)

	time.Sleep(10 * time.Second)

	fmt.Println("push any key!")
	var s string
	fmt.Scanf("%s\n", &s)
}

func go1(r chan Request) {
	t := beginTransaction(r)

	set(t, 1, "foo")
	set(t, 2, "bar")
	set(t, 3, "hoge")

	fmt.Printf("go1 %s\n", get(t, 1))
	fmt.Printf("go1 %s\n", get(t, 2))
	fmt.Printf("go1 %s\n", get(t, 3))

	time.Sleep(2 * time.Second)

	fmt.Printf("go1 %s\n", get(t, 1))
	fmt.Printf("go1 %s\n", get(t, 2))
	fmt.Printf("go1 %s\n", get(t, 3))

	endTransaction(t)

	time.Sleep(5 * time.Second)

	t = beginTransaction(r)

	fmt.Printf("go1 %s\n", get(t, 1))
	fmt.Printf("go1 %s\n", get(t, 2))
	fmt.Printf("go1 %s\n", get(t, 3))

	endTransaction(t)
}

func go2(r chan Request) {

	time.Sleep(1 * time.Second)

	t := beginTransaction(r)

	s := get(t, 2)

	fmt.Printf("go2 %s\n", s)

	set(t, 2, "あああああ")

	time.Sleep(3 * time.Second)

	endTransaction(t)
}
