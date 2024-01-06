package main

import (
	"fmt"
)

// Реализация через канал с пустой структурой
func main() {
	c := make(chan string)
	str := make(chan struct{})
	go func() {
		for {
			select {
			case <-str:
				c <- "Goroutine is closed"
				return
			}
		}
	}()
	str <- struct{}{}
	fmt.Print(<-c)
}
