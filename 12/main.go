package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, s1 := range s {
		set[s1] = struct{}{}
	}
	for k := range set {
		fmt.Print(k, " ")
	}
}
