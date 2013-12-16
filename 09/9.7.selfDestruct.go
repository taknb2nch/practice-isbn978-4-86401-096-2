package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan bool)
	count := make(chan int)

	go cancel(abort)
	go countDown(count)

	for {
		select {
		case i := <-count:
			if i == 0 {
				selfDestruct()
				return
			}

			fmt.Printf("%d seconds remaing\n", i)

		case a := <-abort:
			if a {
				fmt.Println("Self Destruct aborted\n")
			} else {
				selfDestruct()
			}

			return
		}
	}
}

func cancel(reply chan bool) {
	fmt.Println("This program will self destruct, do you wish to cancel\n")

	var r int

	fmt.Scanf("%c\n", &r)

	switch r {
	case 'y':
		reply <- true
	case 'Y':
		reply <- true
	default:
		reply <- false
	}
}

func countDown(count chan int) {
	for i := 10; i >= 0; i-- {
		count <- 1
		time.Sleep(time.Second)
	}
}

func selfDestruct() {
	fmt.Printf("Self destruct feature not yet implemented\n")
}
