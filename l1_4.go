package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ctx context.Context, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: Channel closed, exiting\n", id)
				return
			}
			fmt.Printf("Worker %d: Processed job %d\n", id, job)
		case <-ctx.Done():
			fmt.Printf("Worker %d: Received shutdown signal, exiting\n", id)
			return
		}
	}
}

func main() {
	var workerCount int
	fmt.Println("Enter number of workers")
	fmt.Scan(&workerCount)

	jobs := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, ctx, &wg, jobs)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				job := rand.Intn(100)
				jobs <- job
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived shutdown signal, cancelling context...")

	cancel()

	wg.Wait()

	fmt.Println("All workers have exited. Program terminated.")
}
