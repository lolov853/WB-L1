package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		return index
	}

	return -1
}

func main() {
	arr := []int{3, 8, 15, 23, 42, 56, 78}
	target := 23

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
