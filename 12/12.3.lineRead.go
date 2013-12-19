package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("11.3.lineRead.go")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	lineNumber := 1
	lineReader := bufio.NewReaderSize(file, 20)

	for line, isPrefix, e := lineReader.ReadLine(); e == nil; line, isPrefix, e = lineReader.ReadLine() {
		fmt.Printf("%.3d:", lineNumber)

		lineNumber++

		os.Stdout.Write(line)

		if isPrefix {
			fmt.Print("###")
			for {
				line, isPrefix, e = lineReader.ReadLine()

				os.Stdout.Write(line)

				if !isPrefix {
					break
				}
			}
		}

		fmt.Printf("\n")
	}
}
