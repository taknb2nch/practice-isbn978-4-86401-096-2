package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a file name")
	var s string

	fmt.Scanf("%s\n", &s)

	fi, err := os.Stat(s)

	if err != nil {
		fmt.Println("%s dose not exist!\n", s)
		return
	}

	if fi.IsDir() {
		fmt.Printf("%s is a directory\n", s)
	}

	mode := fi.Mode()

	if mode&os.ModeSymlink == os.ModeSymlink {
		fmt.Printf("%s is a symbolic link\n", s)
	}
}
