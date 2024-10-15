package main

import (
	"context"
	"fmt"
	"time"
)

func workerWithChannel(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Worker with channel: received stop signal")
			return
		default:
			fmt.Println("Worker with channel: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with context: stopping...")
			return
		default:
			fmt.Println("Worker with context: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	go workerWithChannel(stopChan)

	go workerWithContext(ctx)

	time.Sleep(3 * time.Second)

	close(stopChan)

	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Main function finished.")
}
