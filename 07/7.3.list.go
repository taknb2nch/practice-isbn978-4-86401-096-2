package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l.PushBack(42)
	l.PushBack(main)
	l.PushBack("A string")

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v\n", e.Value)
	}
}
