package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	fmt.Print("intput number > ")
	fmt.Scanf("%d", &i)

	str := strconv.FormatInt(int64(i), 10)

	hex, _ := strconv.ParseInt(str, 16, 64)

	fmt.Printf("%d\n", hex)

	for i := 0; i < 36; i++ {
		// FormatInt(i int64, base int) string
		fmt.Println(i, strconv.FormatInt(int64(i), 16))
	}

	// ParseInt(s string, base int, bitSize int) (i int64, err error)
	// Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64
	if _, err := strconv.ParseInt("hoge", 16, 32); err != nil {
		fmt.Println(err)
	}

}