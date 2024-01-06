package main

import (
	"fmt"
	"sync"
)

// Реализация с WaitGroup
func main() {
	var sum int
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for _, n := range nums {
		go func(n int) {
			sum += n * n
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println(sum)
}
