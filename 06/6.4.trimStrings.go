package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\tThis is a string \t hogehoe \n"
	str = strings.Trim(str, " \t\n\r")
	words := strings.Split(str, " ")

	for _, word := range words {
		fmt.Printf("%s\n", word)
	}
}
