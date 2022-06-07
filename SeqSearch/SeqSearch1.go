package main

import "fmt"

func sequentialSearch(n int, arr [5]int, target int) int { //unsorted

	for i := 0; i < n; i++ {

		if arr[i] == target {
			return i
		}

	}
	return -1
}

func main() {
	const size int = 5
	arr := [size]int{10, 30, 20, 8, 5}
	var target int = 20
	result := sequentialSearch(size, arr, target)

	if result == -1 {
		fmt.Printf("Target %+v is not found", target)

	} else {
		fmt.Printf("Target %+v is found at position %+v", target, result+1)
	}
}
