package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (s *SafeMap) Write(key, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
	fmt.Printf("Wrote: [%d] = %d\n", key, value)
}

func main() {
	safeMap := SafeMap{
		m: make(map[int]int),
	}

	var wg sync.WaitGroup
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				safeMap.Write(i*10+j, i*10+j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Final map:", safeMap.m)
}
