package main

func remove(ar []any, i int) []any {
	if i >= len(ar) || i < 0 {
		return nil
	}
	return append(ar[:i], ar[i+1:]...)
}
