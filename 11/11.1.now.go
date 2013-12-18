package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("%d seconds since the Epoc\n", now.Unix())
	fmt.Printf("%d nanoseconds since the Epoc\n", now.UnixNano())

	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	fmt.Println(now.Clock()) // 時間
	fmt.Println(now.Date())  // 日付

	fmt.Println(now.ISOWeek())
	fmt.Println(now.Local())
	fmt.Println(now.Location())

	fmt.Println(now.String())
	fmt.Println(now.UTC())
}
