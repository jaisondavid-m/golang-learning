package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	select {
	case <-time.After(2*time.Second):
		fmt.Println("Operation Completed")
	case <-ctx.Done():
		fmt.Println("Operation Cancelled:",ctx.Err())
	}

}