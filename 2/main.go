package main

import (
	"fmt"
)

func main() {
	b := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	for _, n := range nums {
		go func() {
			c := <-b
			fmt.Println(c * c)
		}()
		b <- n
	}
	fmt.Scanln()
}

//Реализация с WaitGroup
// func main() {
// 	nums := []int{2, 4, 6, 8, 10}
// 	wg := new(sync.WaitGroup)
// 	wg.Add(5)
// 	for _, num := range nums {
// 		go func(num int) {
// 			fmt.Println(num * num)
// 			wg.Done()
// 		}(num)
// 	}
// 	wg.Wait()
// }
