package main

import (
	"fmt"
)

func main() {

	ch := make(chan int,3)

	ch <- 100
	ch <- 200
	ch <- 300

	close(ch)

	for value := range ch {
		fmt.Println(value)
	}

}