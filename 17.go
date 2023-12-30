package main

func binarySearch(ar []int, s int) int {
	var left int
	right := len(ar) - 1
	for left <= right {
		mid := int((left + right) / 2)
		value := ar[mid]
		if value == s {
			return mid
		} else if value < s {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
