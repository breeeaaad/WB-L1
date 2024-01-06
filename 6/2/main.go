package main

import (
	"context"
	"fmt"
	"sync"
)

// Реализация через контекст
func main() {
	ctx, close := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): //Под капотом пустой канал типа структуры
				fmt.Println("Goroutine is closed")
				return
			}
		}
	}()
	close()
	wg.Wait() //WaitGroup ждет выполнения горутины
}
