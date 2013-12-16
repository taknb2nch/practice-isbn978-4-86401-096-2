package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[2] = "First Value"

	c := make(chan bool, 1)

	go func(ch chan<- bool) {
		m[2] = "Second Value"
		ch <- true
	}(c)

	<-c

	fmt.Printf("%s\n", m[2])
}
