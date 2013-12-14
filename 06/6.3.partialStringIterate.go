package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	f1()
	f2()
}

func f1() {
	str := "あいうえお"
	rune := make([]byte, 0, 4)

	for i := 0; i < len(str); i++ {
		rune = append(rune, str[i])

		if utf8.FullRune(rune) {
			char, _ := utf8.DecodeRune(rune)
			fmt.Printf("%c", char)
			rune = rune[0:0]
		}
	}

	fmt.Println()
}

func f2() {
	str := "あいうえお"
	bytes := str[0:7]
	str2 := string(bytes)

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%X\n", str2[i])
	}

	for i, c := range str2 {
		if c == 0XFFFD {
			fmt.Printf("%c", c)
			str2 = str2[i:]
			break
		} else {
			fmt.Printf("%c", c)
		}
	}

	fmt.Println()
}
