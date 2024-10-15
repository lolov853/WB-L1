package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- num * int(num)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, ch)
	}

	wg.Wait()
	close(ch)
	var d int = 0
	for num := range ch {
		d += num
	}
	fmt.Println(d)

}
