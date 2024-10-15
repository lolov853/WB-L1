package main

import (
	"fmt"
	"strings"
)

func hasUniqueCharacters(s string) bool {
	s = strings.ToLower(s)
	charSet := make(map[rune]bool)

	for _, char := range s {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

func main() {
	fmt.Println(hasUniqueCharacters("abcd"))
	fmt.Println(hasUniqueCharacters("abCdefAaf"))
	fmt.Println(hasUniqueCharacters("aabcd"))
	fmt.Println(hasUniqueCharacters("Unique"))
}
