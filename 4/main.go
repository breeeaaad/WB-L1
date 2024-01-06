package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

// Для завершения программы использую Graceful Shutdown
func main() {
	var N int
	fmt.Scanln(&N)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		fmt.Print("Stop")
		stop()
		time.Sleep(time.Minute)
	}()
	r := make(chan int)
	for i := 0; i < N; i++ {
		go worker(r)
		r <- rand.Intn(10)
	}
	<-ctx.Done() //На этом этапе главный поток блокируется и ждет Ctrl+C
}

func worker(r chan int) {
	fmt.Println(<-r)
}
