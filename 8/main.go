package main

import "fmt"

//Использование побитовых операций
func setBit(num int64, pos, val int) int64 {
	if val == 1 {
		return num | (1 << pos)
	}
	return num &^ (1 << pos)
}

func main() {
	fmt.Println(setBit(2, 1, 0))
}
