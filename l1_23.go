package main

import (
	"fmt"
)

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	i := 3

	slice = removeElement(slice, i)

	fmt.Println(slice)
}
