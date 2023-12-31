package main

import "fmt"

func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	pivot := ar[0]
	var less, greater []int
	for _, num := range ar[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func main() {
	fmt.Print(quickSort([]int{32, 99, 23, -24}))
}
