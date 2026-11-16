package main

import "fmt"

func worker(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {

	ch := make(chan int)

	go worker(ch)

	for value := range ch {
		fmt.Println("Received:",value)
	}

	fmt.Println("All values received!")

}