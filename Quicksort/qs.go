package main

import (
	"fmt"
)

func quickSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	mid := (left + right) / 2
	a[mid], a[right] = a[right], a[mid]

	for i, _ := range a {

		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func main() {

	arr := []int{1, 68, -7, 1008, 9, -1000, -100, 230, 506}

	fmt.Println("Unsorted!")
	fmt.Println(arr)
	fmt.Println("Sorted!")
	quickSort(arr)
	fmt.Println(arr)

}
