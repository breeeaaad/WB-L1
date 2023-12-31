package main

import "fmt"

func main() {
	m := make(map[int]int)
	f := []int{1, 2, 3, 4, 5, 7}
	s := []int{5, 4, 1, 3, 6, 12}
	for _, n := range f {
		m[n]++
	}
	for _, n := range s {
		m[n]++
	}
	for i, j := range m {
		if j == 2 {
			fmt.Print(i, " ")
		}
	}
}

// func main() {
// 	f := []int{1, 2, 3, 4, 5, 7}
// 	s := []int{5, 4, 1, 3, 6, 12}
// 	for _, n := range f {
// 		for _, b := range s {
// 			if n == b {
// 				fmt.Print(n, " ")
// 			}
// 		}
// 	}
// }
