package main

import (
	"fmt"
	"sync"
)

func main() {
	var b, c int
	fmt.Scanln(&c)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			b++
			mu.Unlock()
		}(wg, mu)
	}
	wg.Wait()
	fmt.Println(b)
}
