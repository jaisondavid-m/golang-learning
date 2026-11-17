package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Before cancellation:",ctx.Err())

	cancel()

	fmt.Println("After Cancellation:",ctx.Err())

}