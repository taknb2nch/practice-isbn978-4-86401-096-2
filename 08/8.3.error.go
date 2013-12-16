package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(i int) (float64, error) {
	if i < 0 {
		//
		return 0, errors.New("無効な引数を指定しました。")
	}

	return math.Sqrt(float64(i)), nil
}

func main() {
	r, _ := sqrt(2)

	fmt.Printf("sqrt(2) = %f\n", r)

	fmt.Println("Enter another number\n")

	var i int
	fmt.Scanf("%d\n", &i)

	if root, err := sqrt(i); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("sqrt(%d) = %f\n", i, root)
	}
}
