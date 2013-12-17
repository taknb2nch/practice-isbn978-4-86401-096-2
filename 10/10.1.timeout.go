package main

import (
	"fmt"
	"time"
)

func readString(reply chan string) {
	fmt.Println("Enter some text")

	var s string
	fmt.Scanf("%s\n", &s)

	reply <- s
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}

func main() {
	t := make(chan bool)
	s := make(chan string)

	go readString(s)
	go timeout(t)

	select {
	case msg := <-s:
		fmt.Printf("Received: %s\n", msg)
	case <-t:
		fmt.Println("Timeout")
	}
}
