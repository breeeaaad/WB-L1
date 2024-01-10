package main

import "fmt"

//Конвеер чисел
func main() {
	nums := []int{5, 2, 5, 23, 6}
	x := make(chan int, 5)
	x2 := make(chan int, 5)
	for _, num := range nums {
		x <- num
	}
	close(x)
	for d := range x {
		x2 <- d * 2
	}
	close(x2)
	for n := range x2 {
		fmt.Println(n)
	}
}
