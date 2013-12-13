package main

import (
	"fmt"
)

func main() {
	//s0 := make([]int, 2, 10)
	s0 := make([]int, 2, 3)
	s1 := append(s0, 2)
	// appendするときにcapを超えるので新しい配列を作成し、s1をコピーしてから追加される。
	s3 := append(s1, 4)
	s2 := append(s0, 3)
	//s3 := append(s1, 4)

	fmt.Printf("Element: %d %d\n", s1[2], s2[2])
	fmt.Println(s0, s1, s2, s3)
	fmt.Printf("len=%d, cap=%d\n", len(s3), cap(s3))

	s0 = []int{0, 1}
	s1 = append(s0, 2)
	s2 = append(s0, 3)

	fmt.Printf("Element: %d %d\n", s1[2], s2[2])
	fmt.Println(s0, s1, s2)
}
