package main

import (
	"fmt"
	"sync"
)

type LazyInit struct {
	once  sync.Once
	value int
}

func (l *LazyInit) Value() int {
	l.init()
	return l.value
}

func (l *LazyInit) init() {
	l.once.Do(func() {
		l.value = 42
	})
}

func (l *LazyInit) SetValue(v int) {
	l.value = v
}

func main() {
	var l LazyInit
	fmt.Printf("%d\n", l.Value())

	l.SetValue(12)
	fmt.Printf("%d\n", l.Value())
}
