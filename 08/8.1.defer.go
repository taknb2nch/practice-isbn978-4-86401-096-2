package main

import (
	"fmt"
	"sync"
)

func main() {
	var broken func()
	var lock sync.Mutex

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	callLocked(&lock, broken)
}

func callLocked(lock *sync.Mutex, f func()) {
	lock.Lock()

	defer lock.Unlock()

	f()
}
