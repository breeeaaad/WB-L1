package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var TTL int64
	fmt.Scanln(&TTL)
	now := time.Now()
	info := make(chan int)
	for time.Since(now) < time.Second*time.Duration(TTL) {
		go func() {
			fmt.Println(<-info)
		}()
		info <- rand.Intn(10)
	}
}
