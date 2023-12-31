package main

import (
	"fmt"
	"time"
)

//Можно через канал или через context
//Канал можно закрыть или передать boolean тип и проверить
//Можно отправить пустую структуру и, если канал не пустой, закрыть

// func main() {
// 	ctx, close := context.WithCancel(context.Background())
// 	wg := new(sync.WaitGroup)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done(): //Под капотом пустой канал типа структуры
// 				fmt.Println("Goroutine is closed")
// 				return
// 			}
// 		}
// 	}()
// 	time.Sleep(time.Second)
// 	close()
// 	wg.Wait() //WaitGroup ждет выполнения горутины
// }

func main() {
	c := make(chan string)
	str := make(chan struct{})
	go func() {
		for {
			select {
			case <-str: //Метод с отправкой пустой структуры(также можно и boolean тип проверить)
				c <- "Goroutine is closed" //можно закрыть канал и проверить с помощью if _,is_open:=<-c
				return
			}
		}
	}()
	time.Sleep(time.Second)
	str <- struct{}{}
	fmt.Print(<-c) //Второй способ ждать выполнения горутины
}
