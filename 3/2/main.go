package main

import (
	"fmt"
	"sync"
)

// Реализация с WaitGroup и Мутексом для безопасности итераций
func main() {
	var sum int
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(5)
	for _, n := range nums {
		go func(n int) {
			mu.Lock()
			sum += n * n
			mu.Unlock()
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println(sum)
}
