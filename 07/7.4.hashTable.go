package main

import (
	"container/list"
	"fmt"
)

type Hashable interface {
	Hash() int
	IsEqual(Hashable) bool
}

type HashTable struct {
	table map[int]*list.List
}

func (h HashTable) Find(value Hashable) Hashable {
	l := h.table[value.Hash()]

	if l == nil {
		return nil
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if value.IsEqual(e.Value.(Hashable)) {
			return e.Value.(Hashable)
		}
	}

	return nil
}

func (h *HashTable) Add(value Hashable) {
	if h.Find(value) != nil {
		return
	}

	if h.table == nil {
		h.table = make(map[int]*list.List)
	}

	l := h.table[value.Hash()]

	if l == nil {
		l = new(list.List)
	}

	l.PushBack(value)
	h.table[value.Hash()] = l
}

// 独自メソッドを持たせてHashableとして扱うためにtypeで定義する
type str string

func (s str) Hash() int {
	return len(s)
}

func (s str) IsEqual(other Hashable) bool {
	return s == other.(str)
}

func main() {
	var h HashTable

	h.Add(str("Foo"))
	h.Add(str("Foo"))
	h.Add(str("Bar"))
	h.Add(str("Wibble"))

	fmt.Printf("%v %v %v\n", h.Find(str("Foo")), h.Find(str("Bar")), h.Find(str("Wibble")))
}
