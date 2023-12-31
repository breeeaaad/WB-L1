package main

import "fmt"

func remove(arr []int, i int) []int {
	if i >= len(arr) || i < 0 {
		return nil
	}
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(remove(nums, 3))
}
