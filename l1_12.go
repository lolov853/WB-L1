package main

import "fmt"

func uniqueSet(strings []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := uniqueSet(strings)
	fmt.Println(set)
}
