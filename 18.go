package main

import (
	"fmt"
	"sync"
)

func inc(n int) {
	var b int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i <= n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			b++
			mu.Unlock()
		}(wg, mu)
	}
	wg.Wait()
	fmt.Print(b)
}
