package main

import (
	"fmt"
	"time"
)

func writeToFirstChannel(arr []int, ch1 chan int) {
	for _, num := range arr {
		ch1 <- num
	}
	close(ch1)
}

func processNumbers(ch1, ch2 chan int) {
	for num := range ch1 {
		ch2 <- num * 2
	}
	close(ch2)
}

func printResults(ch2 chan int) {
	for result := range ch2 {
		fmt.Println("Result:", result)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go writeToFirstChannel(arr, ch1)

	go processNumbers(ch1, ch2)

	printResults(ch2)

	time.Sleep(1 * time.Second)
}
