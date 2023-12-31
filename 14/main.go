package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := []any{1, "", true, make(chan int)}
	for _, ty := range t {
		fmt.Print(reflect.TypeOf(ty)) //можно еще if value,ok:=ty.(int);if ok ..., но кода много
	}
}
