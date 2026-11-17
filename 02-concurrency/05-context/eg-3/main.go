package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context,wg *sync.WaitGroup) {

	defer wg.Done()

	ticker := time.NewTicker(time.Second*2)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:",ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Processing task...")
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	
	wg.Add(1)

	go worker(ctx,&wg)

	time.Sleep(time.Second*5)

	fmt.Println("Sending cancellation signal...")
	cancel()

	wg.Wait()
	fmt.Println("Application shutdown complete")

}