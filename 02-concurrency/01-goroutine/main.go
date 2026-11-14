package main

import (
	"fmt"
	"sync"
)

func routineWorkers(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Worker id is", id)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go routineWorkers(&wg, 1)
	go routineWorkers(&wg, 2)
	go routineWorkers(&wg, 3)

	wg.Wait()

	fmt.Println("All workers completed!")

}
