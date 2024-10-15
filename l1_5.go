package main

import (
	"fmt"
	"time"
)

func sendValue(ch chan int, done chan struct{}) {
	counter := 0
	for {
		select {
		case ch <- counter:
			counter++
			time.Sleep(500 * time.Millisecond)

		case <-done:
			fmt.Println("Finished sending")
			return
		}
	}

}
func recieveValue(ch chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
func main() {
	chn := make(chan int)
	done := make(chan struct{})

	var duration int
	fmt.Println("Enter duration: ")
	fmt.Scan(&duration)

	go sendValue(chn, done)

	go recieveValue(chn)

	time.Sleep(time.Duration(duration) * time.Second)

	close(chn)

	time.Sleep(1 * time.Second)
	fmt.Println("Program finished.")
}
