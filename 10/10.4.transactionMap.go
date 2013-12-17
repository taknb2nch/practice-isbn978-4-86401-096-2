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

	// set(m, 1, "foo")
	// set(m, 2, "bar")
	// set(m, 3, "hoge")

	// t := beginTransaction(m)

	// // set(t, 1, "foo")
	// // set(t, 2, "bar")
	// // set(t, 3, "hoge")

	// s := get(t, 2)

	// fmt.Println(s)

	// set(t, 2, "あああああ")

	// endTransaction(t)

	// fmt.Printf("Set %s\n", get(m, 1))
	// fmt.Printf("Set %s\n", get(m, 2))

	// fmt.Printf("Set %s\n", get(m, 3))

	go go1(m)
	go go2(m)

	time.Sleep(10 * time.Second)

}

func go1(r chan Request) {
	set(r, 1, "foo")
	set(r, 2, "bar")
	set(r, 3, "hoge")

	fmt.Printf("go1 %s\n", get(r, 1))
	fmt.Printf("go1 %s\n", get(r, 2))
	fmt.Printf("go1 %s\n", get(r, 3))

	time.Sleep(2 * time.Second)

	fmt.Printf("go1 %s\n", get(r, 1))
	fmt.Printf("go1 %s\n", get(r, 2))
	fmt.Printf("go1 %s\n", get(r, 3))

	time.Sleep(5 * time.Second)

	fmt.Printf("go1 %s\n", get(r, 1))
	fmt.Printf("go1 %s\n", get(r, 2))
	fmt.Printf("go1 %s\n", get(r, 3))
}

func go2(r chan Request) {

	time.Sleep(1 * time.Second)

	t := beginTransaction(r)

	s := get(t, 2)

	fmt.Printf("go2 %s\n", s)

	set(t, 2, "あああああ")

	time.Sleep(5 * time.Second)

	endTransaction(t)
}
