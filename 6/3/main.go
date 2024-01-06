package main

import "fmt"

//Через закрытие канала и проверку на закрытие
func main() {
	c := make(chan struct{})
	info := make(chan string)
	go func() {
		for {
			if _, is_open := <-c; !is_open {
				info <- "Chan is closed"
			}
		}
	}()
	close(c)
	fmt.Print(<-info)
}
