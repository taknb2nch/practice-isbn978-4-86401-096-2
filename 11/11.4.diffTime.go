package main

import (
	"fmt"
	"time"
)

func main() {
	parsed, _ := time.Parse("2006/01/02", "1982/06/15")

	now := time.Now()

	parsedSeconds := parsed.Unix()

	fmt.Printf("%d seconds diffrence\n", now.Unix()-parsedSeconds)

	diff := now.Sub(parsed)
	fmt.Printf("%s difference\n", diff.String())
}
