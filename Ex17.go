package main

import (
	"fmt"
	"sort"
)

func search(key int, arr []int) (pos int) {
	mid := len(arr) / 2
	switch {
	case len(arr) == 0:
		pos = -1
	case arr[mid] > key:
		pos = search(key, arr[:mid])
	case arr[mid] < key:
		pos = search(key, arr[mid+1:])
		if pos >= 0 {
			pos += mid + 1
		}
	default:
		pos = mid
	}
	return
}

func Ex17() {
	arr := []int{9, 12, 5, 6, 4, 2, 1}
	arr = QuickSortStart(arr)
	key := 9

	keyPos := search(key, arr)
	fmt.Printf("index of key using hand made search: %d, ", keyPos)

	keyPos = sort.Search(len(arr), func(i int) bool {
		return key == arr[i]
	})
	fmt.Printf("index of key using standard library sort.Search: %d", keyPos)

}
