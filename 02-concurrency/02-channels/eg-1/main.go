package main

import "fmt"

func calculate(ch chan int,a,b int) {
	result := a + b
	ch <- result
}

func main() {
	a := 5
	b := 10
	ch := make(chan int, 2)
	go calculate(ch,a,b)
	result := <-ch
	fmt.Println("Result:",result)
}