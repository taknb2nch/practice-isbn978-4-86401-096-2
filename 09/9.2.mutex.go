package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cur := runtime.GOMAXPROCS(2)
	fmt.Println(cur)

	m := make(map[int]string)
	m[2] = "first Value"

	var lock sync.Mutex

	go func() {
		lock.Lock()
		m[2] = "Second Value"
		lock.Unlock()
	}()

	// 一瞬waitを入れるとgoroutineで先にロックされる
	time.Sleep(1 * time.Nanosecond)

	lock.Lock()
	v := m[2]
	lock.Unlock()

	time.Sleep(1 * time.Second)

	fmt.Printf("%s\n", v)
}
