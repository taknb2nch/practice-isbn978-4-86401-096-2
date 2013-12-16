package main

import (
	"fmt"
)

func badFunction() {
	fmt.Println("Select Panic type (0=no panic, 1=int, 2=runtime panic)")

	var choice int
	fmt.Scanf("%d\n", &choice)

	switch choice {
	case 1:
		panic(0)
	case 2:
		var invalid func()
		invalid()
	}
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			default:
				panic(x)
			case int:
				fmt.Printf("Function paniced with a very unhelpful error: %d\n", x)
			}
		}
	}()

	badFunction()

	fmt.Println("Program exited normally")
}
