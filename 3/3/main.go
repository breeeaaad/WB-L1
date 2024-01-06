package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализация с WaitGroup и atomic для той же безопасности + скорость так как функция реализована на ассемблере
func main() {
	var sum int64
	nums := []int64{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for _, n := range nums {
		go func(n int64) {
			atomic.AddInt64(&sum, n*n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println(sum)
}
