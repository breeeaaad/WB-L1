package main

import "fmt"

//O(n+m) по скорости
func main() {
	m := make(map[int]struct{})
	f := []int{1, 2, 3, 4, 5, 7}
	s := []int{5, 4, 1, 3, 6, 12}
	for _, n := range f {
		m[n] = struct{}{}
	}
	for _, n := range s {
		if _, ok := m[n]; ok {
			fmt.Print(n, " ")
		}
	}

}

// O(n*m) по скорости, поэтому закоментил
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
