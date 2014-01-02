package main

import (
	"fmt"
)

func main() {
	s := make(chan *string)

	go func() {
		var buffer string

		fmt.Scanf("%s\n", &buffer)

		s <- &buffer
	}()

	fmt.Printf("%s\n", *(<-s))

	// s := make(chan string)

	// go func() {
	// 	var buffer string

	// 	fmt.Scanf("%s\n", &buffer)

	// 	s <- buffer
	// }()

	// fmt.Printf("%s\n", <-s)
}
