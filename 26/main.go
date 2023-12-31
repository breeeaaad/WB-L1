package main

import (
	"fmt"
	"strings"
)

//set можно создать с помощью map со значениями типа bool или struct
//struct не использует память

func is_unique(s string) bool {
	m := make(map[string]struct{})
	for _, c := range strings.ToLower(s) {
		_, ok := m[string(c)]
		if ok {
			return false
		}
		m[string(c)] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(is_unique("abcd"))
	fmt.Println(is_unique("abCdefAaf"))
	fmt.Print(is_unique("aabcd"))
}
