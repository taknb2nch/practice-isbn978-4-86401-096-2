package main

import (
	"fmt"
)

func main() {

	s1 := []int{1, 2, 3}
	s2 := truncate(s1)
	s2[0] = 10

	fmt.Println(s1, s2)

}

func truncate(slice []int) []int {
	var s []int = make([]int, len(slice))
	copy(s, slice)
	return s
}
