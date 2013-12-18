package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Timer expired")
	})

	timer := time.NewTimer(3 * time.Second)

	fmt.Printf("Current time: %d nanoseconds\n", <-timer.C)

	// 両方のタイマーが同時に動いている？
}
