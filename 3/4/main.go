package main

import "fmt"

// Реализация с каналами
func main() {
	var sum int
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)
	for _, n := range nums {
		go func(n int) {
			ch <- n * n
		}(n)
	}
	for i := 0; i < 5; i++ {
		sum += <-ch
	}
	fmt.Println(sum)
}
