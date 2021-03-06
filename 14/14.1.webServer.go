package main

import (
	"fmt"
	"net/http"
)

type webCounter struct {
	count chan int
}

func NewCounter() *webCounter {
	counter := new(webCounter)
	counter.count = make(chan int, 1)

	go func() {
		for i := 1; ; i++ {
			counter.count <- i
		}
	}()

	return counter
}

func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
	if rq.URL.Path != "/" {
		r.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(r, "You are visiter %d", <-w.count)
}

func main() {
	err := http.ListenAndServe(":8000", NewCounter())

	if err != nil {
		fmt.Println("Server failed.", err)
	}
}
