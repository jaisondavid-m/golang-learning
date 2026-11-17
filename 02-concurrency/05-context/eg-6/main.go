package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	deadline := time.Now().Add(time.Second*5)

	ctx, cancel := context.WithDeadline(
		context.Background(),
		deadline,
	)

	defer cancel()

	select {
	case <-time.After(5*time.Second):
		fmt.Println("Operation Completed")
	case <-ctx.Done():
		fmt.Println("Operation Cancelled:",ctx.Err())
	}

}