package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Printed in the background.")

	i := 1

	// 現時点でのiの値のコピーが渡されます。
	go fmt.Printf("Currently, i is %d\n", i)

	// 関数外のiを直接参照しているのでiの値がどうなるかは分かりません。
	go func() {
		fmt.Printf("i: %d\n", i)
	}()

	i++

	time.Sleep(5 * time.Second)
}
