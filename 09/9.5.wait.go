package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(5))

	var w sync.WaitGroup

	for _, v := range os.Args {
		w.Add(1)

		go func(str string) {
			fmt.Printf("%s\n", strings.ToUpper(str))
			w.Done()
		}(v)
	}

	w.Wait()
}
