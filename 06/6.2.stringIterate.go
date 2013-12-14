package main

import (
	"fmt"
)

func main() {
	str := "あいうえお"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println()

	for _, c := range str {
		fmt.Printf("%c", c)
	}

	fmt.Println()
}
