package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) Whoami() {
	fmt.Print("I am ", h.Name)
}

type Action struct {
	Human //Для *наследования* пишем лишь название родительской структуры
}

func main() {
	var a Action
	a.Name = "Mike"
	a.Whoami()
}
