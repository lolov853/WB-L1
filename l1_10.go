package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(math.Floor(temp/10.0)) * 10
		tempGroups[group] = append(tempGroups[group], temp)
	}

	for group, temps := range tempGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
