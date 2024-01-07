package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var N int
	fmt.Scanln(&N)
	r := make(chan int)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT)
	for i := 0; i < N; i++ {
		go worker(i, r)
	}
	for {
		select {
		case <-s:
			fmt.Print("Handle Ctrl+C")
			close(r) //Закрытие канала, чтобы больше не передавать значения
			return
		default:
			r <- rand.Intn(10)
		}
	}
}

func worker(id int, r chan int) {
	for n := range r { //Если канал закрыт, цикл прекратиться
		fmt.Println(id, "-", n)
	}
}
