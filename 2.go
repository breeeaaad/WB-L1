package main

import (
	"fmt"
)

func printsqr() {
	b := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	fmt.Print("2. ")
	for _, n := range nums {
		go func() {
			c := <-b
			fmt.Println(c * c)
		}()
		b <- n
	}
}
