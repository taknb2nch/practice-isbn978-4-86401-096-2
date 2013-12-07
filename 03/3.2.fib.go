package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Printf("Conpute how many Fibonacci numbers?")
	fmt.Scanf("%d", &n)

	last := big.NewInt(1)
	current := big.NewInt(1)

	for i := 0; (i < n) && (i < 2); i++ {
		fmt.Printf("1\n")
	}

	for i := 2; i < n; i++ {
		// a.Add(b, c) -> a = b + c
		last.Add(last, current)
		last, current = current, last

		fmt.Printf("%s\n", current.String())
	}

	//
	a, b := 1, 2

	fmt.Println(a, b)
	// Output:1 2

	// 左から順次代入されるのではなく、一気に代入される
	a, b = b, a

	fmt.Println(a, b)
	// Output:2 1

	f1 := 0.0
	d1, d2 := new(big.Rat), new(big.Rat)

	// d1.SetString("0.0")
	// d2.SetString("0.1")
	d1.SetFloat64(0.0)
	d2.SetFloat64(0.1)

	for i := 0; i < 10; i++ {
		f1 += 0.1
		d1.Add(d1, d2)
	}

	df, _ := d1.Float64()

	fmt.Printf("%f, %f\n", f1, df)
	fmt.Printf("%1.15e, %1.15e\n", f1, df)
}