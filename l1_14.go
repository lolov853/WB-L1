package main

import (
	"fmt"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is channel")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var x interface{}

	x = 42
	determineType(x)

	x = "hello"
	determineType(x)

	x = true
	determineType(x)

	x = make(chan int)
	determineType(x)
}
