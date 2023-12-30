package main

import "fmt"

type Human struct{}

func (Human) Whoami() {
	fmt.Println("1. I am human")
}

type Action struct {
	Human //Для *наследования* пишем лишь название родительской структуры
}
