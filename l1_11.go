package main

import "fmt"

func intersection(set1, set2 []int) []int {
	result := []int{}
	setMap := make(map[int]bool)

	for _, v := range set1 {
		setMap[v] = true
	}

	for _, v := range set2 {
		if setMap[v] {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println(result)
}
