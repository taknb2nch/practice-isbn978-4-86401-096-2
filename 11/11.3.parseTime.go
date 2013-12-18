package main

import (
	"fmt"
	"time"
)

func main() {
	var t string

	fmt.Println("Enter a time")
	fmt.Scanf("%s\n", &t)

	//parsed, err := time.Parse("03:04PM", t)
	//parsed, err := time.Parse("2006-01-02 15:04:05 -0700", t)
	parsed, err := time.Parse("2006-01-02T15:04:05", t)
	//parsed, err := time.Parse("2006/01/02 15:04:05", "2013/12/18 14:11:30")

	if err != nil {
		fmt.Println(err)
		parsed, err = time.Parse("15:04", t)
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Time in seconds since the Epoc: %d\n", parsed.Unix())
	}

	fmt.Printf("%s\n", parsed.Format(time.RFC3339))
	fmt.Println(parsed.Format("2006/01/02 15:04:05 MST"))
}
