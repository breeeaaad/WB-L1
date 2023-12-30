package main

import "fmt"

func rev(a, b int) {
	a, b = b, a //первый способ реверса
	a ^= b      //Второй способ
	b ^= a
	a ^= b
	fmt.Print("13. ")
	fmt.Println(a)
	fmt.Println(b)
}
