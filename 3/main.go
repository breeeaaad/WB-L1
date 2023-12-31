package main

import (
	"fmt"
)

func main() {
	var sum int
	b := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	for _, n := range nums {
		go func() {
			c := <-b
			sum += c * c
		}()
		b <- n
	}
	fmt.Print(sum)
}

// func main() {
// 	var sum int
// 	nums := []int{2, 4, 6, 8, 10}
// 	wg := new(sync.WaitGroup)
// 	wg.Add(5)
// 	for _, n := range nums {
// 		go func(n int) {
// 			sum += n * n
// 			wg.Done()
// 		}(n)
// 	}
// 	wg.Wait()
// 	fmt.Println(sum)
// }
