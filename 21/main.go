package main

import "fmt"

type USB interface {
	connectWithUsbCable()
}

type MemoryCard struct{}

func (m MemoryCard) insert() {
	fmt.Println("Карта памяти успешно вставлена!")
}

//В других языках нет такого *наследования и имплементирования*, но можно ли сделать так чтобы методы и поля MemoryCard вызывались напрямую?
type CardReader struct {
	//MemoryCard // имею ввиду вот такую реализацию
	memoryCard MemoryCard
}

func (c CardReader) connectWithUsbCable() {
	c.memoryCard.insert()
}

func main() {
	c := CardReader{}
	c.connectWithUsbCable()
}
