package main

import (
	"fmt"
)

type stackEntry struct {
	next  *stackEntry
	value interface{}
}

type stack struct {
	top *stackEntry
}

func (s *stack) Push(v interface{}) {
	var e stackEntry
	e.value = v
	e.next = s.top
	s.top = &e
}

func (s *stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	v := s.top.value
	s.top = s.top.next

	return v
}

type Stack interface {
	Push(interface{})
	Pop() interface{}
}

// interfaceのインスタンスを作成するメソッドはNew+interface名
func NewStack() Stack {
	return &stack{}
}

func main() {
	st := NewStack()
	st.Push(12)
	st.Push(12)
	st.Push(1)
	st.Push("あいうえお")
	st.Push("abcdefg")

	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop(), st.Pop())
	fmt.Println(st.Pop())
}
