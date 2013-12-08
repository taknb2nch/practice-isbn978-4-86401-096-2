package main

import (
	"fmt"
)

type A struct {
}

type B struct {
}

func (_ A) Print() {
	fmt.Printf("A printed\n")

}

func (_ B) Print() {
	fmt.Printf("B printed\n")

}

func (a A) PrintA() {
	a.Print()
}

type C struct {
	A
	*B
}

func main() {
	var c C
	c.B = &B{}
	c.PrintA()

	// 暗黙の継承のためA.PrintかB.Printか分からない
	//c.Print()

	c.A.Print()
	c.B.Print()
}
