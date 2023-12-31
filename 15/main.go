package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(l int) string {
	var s string
	for l > 0 {
		s += "a"
		l--
	}

	return s
}

func someFunc() {
	v := createHugeString(1 << 10) // 1024 bytes
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Print(justString)
}
