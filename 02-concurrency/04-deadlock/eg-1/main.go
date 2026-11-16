package main

import "fmt"

func main() {

	ch := make(chan int) // user (chan int,1) to prevent to deadlock while executing

	ch <- 100

	fmt.Println(<-ch)

}