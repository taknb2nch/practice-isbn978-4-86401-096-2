package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("11.2.fileRead.go")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	buffer := make([]byte, 100)

	//for n, e := file.Read(buffer); e != nil; n, e = file.Read(buffer) {
	for n, e := file.Read(buffer); ; n, e = file.Read(buffer) {
		if e != nil {
			fmt.Println("#1", e, n)
			break
		}

		if n > 0 {
			// 実際に読み込んだサイズだけをスライスして出力.
			//os.Stdout.Write(buffer[0:n])

			s := string(buffer[0:n])

			fmt.Printf("%s", s)
		} else {
			fmt.Println("#2 n==0")
			break
		}
	}
}
