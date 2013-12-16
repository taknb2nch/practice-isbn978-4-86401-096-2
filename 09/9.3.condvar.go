package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(2))

	m := make(map[int]string)
	m[2] = "First Value"

	var mutex sync.Mutex
	cv := sync.NewCond(&mutex)

	updateCompleted := false

	go func() {
		cv.L.Lock()
		fmt.Println("goroutine lock")

		m[2] = "Second Value"
		updateCompleted = true

		cv.Signal()
		fmt.Println("goroutine signal")
		cv.L.Unlock()
		fmt.Println("goroutine unlock")
	}()

	//time.Sleep(time.Second)

	cv.L.Lock()
	fmt.Println("main lock")

	for !updateCompleted {
		fmt.Println("main wait")
		cv.Wait()
	}

	v := m[2]

	cv.L.Unlock()
	fmt.Println("main unlock")

	fmt.Printf("%s\n", v)
}
