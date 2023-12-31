package main

import (
	"fmt"
	"os"
	"os/signal"
)

func Worker(data chan int) {
	fmt.Println(<-data)
}

func main() {
	var N int
	fmt.Scanln(&N)
	data := make(chan int)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for i := 0; i < N; i++ {
		go Worker(data)
		data <- i
	}
	for {
		select {
		case <-c:
			close(data)
			return
		}
	}
}
