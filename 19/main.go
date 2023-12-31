package main

import "fmt"

func reverse(s string) string {
	runeString := []rune(s)
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}
	return string(runeString)
}

func main() {
	fmt.Print(reverse("главрыба"))
}
