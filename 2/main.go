package main

import (
	"fmt"
	"sync"
)

// Реализация с каналами. Горутина блокируется, пока главный поток не запишет в канал
//Насчет этого метода неуверен
// func main() {
// 	b := make(chan int)
// 	mu := new(sync.Mutex)
// 	nums := []int{2, 4, 6, 8, 10}
// 	for _, n := range nums {
// 		go func() {
// 			c := <-b
// 			fmt.Println(c * c)
// 		}()
// 		b <- n
// 	}
// 	fmt.Scanln()//для выполнения всех горутин
// }

// Реализация с WaitGroup, которая ждет выполнения всех функций
func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for _, num := range nums {
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
