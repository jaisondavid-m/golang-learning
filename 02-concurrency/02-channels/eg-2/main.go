package main

import "fmt"

func main() {
	
	ch := make(chan int,2)  // ,2 for enabling the channal to carry two values

	ch <- 100
	ch <- 200

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}