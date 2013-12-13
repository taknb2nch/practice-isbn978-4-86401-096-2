package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [100]int

	// slow
	for i := 0; i < 10; i++ {
		fmt.Printf("Element %d is %d\n", i, a[i])
	}

	// fast
	subrange := a[1:10]
	//subrange := a[:]
	for i, v := range subrange {
		fmt.Printf("Element %d is %d\n", i, v)
	}

	// parallel
	runtime.GOMAXPROCS(2)
	for i, v := range subrange {
		go fmt.Printf("Element %d is %d\n", i, v)
	}

	time.Sleep(3 * time.Second)

}
