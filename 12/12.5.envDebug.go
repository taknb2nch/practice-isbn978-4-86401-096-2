package main

import (
	"fmt"
	"os"
	"strconv"
)

var debugLevel int

func debugLog(level int, msg string, args ...interface{}) {
	if debugLevel > level {
		fmt.Printf(msg, args...)
	}
}

func main() {
	debug := os.Getenv("DEBUG")

	fmt.Println(debug)

	debugLevel, _ = strconv.Atoi(debug)
	debugLog(3, "Starting\n")
}
