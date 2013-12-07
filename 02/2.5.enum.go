package main

import "fmt"

const (
	Red             = (1 << iota)
	Green           = (1 << iota)
	// 同一行に書いたiotaは同じ値です
	Blue, ColorMask = (1 << iota), (1 << (iota + 1)) - 1
)
// const (
// 	Red       = (1 << iota)
// 	Green     = (1 << iota)
// 	Blue      = (1 << iota)
// 	ColorMask = (1 << iota) - 1
// )

func main() {
	fmt.Println(Red, Green, Blue, ColorMask)
}
