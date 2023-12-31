package main

import "fmt"

func rev(a, b int) {
	a, b = b, a //первый способ реверса
	a ^= b      //второй способ
	b ^= a
	a ^= b
	fmt.Println("13. ", a, "\n", b)
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	a, b = b, a //первый способ
	fmt.Println(a, b)
	b += a //второй способ
	a = b - a
	b -= a
	fmt.Println(a, b)
	a ^= b //третий способ XOR
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
