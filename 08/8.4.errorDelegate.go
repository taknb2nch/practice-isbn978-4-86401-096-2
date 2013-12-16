package main

import (
	"fmt"
	"math"
)

type sqrtError interface {
	invalidArgument(int) (int, error)
}

func sqrt(i int, e sqrtError) (float64, error) {
	for i < 0 {
		var err error
		// if i, err = e.invalidArgument(i); err != nil {
		// 	return 0, err
		// }
		i, err = e.invalidArgument(i)

		if err != nil {
			return 0, err
		}
	}

	return math.Sqrt(float64(i)), nil
}

type sqrtHandler struct{}

func (_ sqrtHandler) invalidArgument(i int) (int, error) {
	fmt.Printf("%d is not valid, Please enter anothor value\n", i)
	// \n を付けていないと後々のScanfで動作がおかしくなります
	fmt.Scanf("%d\n", &i)

	return i, nil
}

func main() {
	fmt.Println("数値を入力してください。")

	var i int
	// \n を付けていないと後々のScanfで動作がおかしくなります
	fmt.Scanf("%d\n", &i)

	if root, err := sqrt(i, sqrtHandler{}); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("sqrt(%d) = %f\n", i, root)
	}
}
