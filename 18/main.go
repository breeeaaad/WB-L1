package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Inc struct {
	//sync.Mutex //Для синхронизации с помощью мьютексов
	c int64
}

func main() {
	var (
		c int
		b Inc
	)
	fmt.Scanln(&c)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			atomic.AddInt64(&b.c, 1) //пакет atomic для счетчика или мьютексы для блокировки изменений
			// b.Lock()
			// b.c++
			// b.Unlock()
		}(wg, mu)
	}
	wg.Wait() //Безопасно вытащить значение при преобразовании с помощью пакета atomic можно с помощью функции atomic.LoadInt64()
	fmt.Println(b.c)
}
