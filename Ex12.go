package main

import "fmt"

func Ex12() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for i := range array {
		set[array[i]] = struct{}{}
	}
	for i := range set {
		fmt.Printf("%s ", i)
	}
}
