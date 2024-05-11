package main

import "fmt"

func QuickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, bottom, top int) []int {
	if bottom < top {
		var p int
		arr, p = partition(arr, bottom, top)
		quickSort(arr, bottom, p-1)
		quickSort(arr, p+1, top)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func Ex16() {
	arr := []int{9, 9, 5, 6, 4, 2, 1}
	arr = QuickSortStart(arr)
	fmt.Println(arr)
}
