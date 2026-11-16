package main

import (
    "fmt"
    "sync"
)

func increment(wg *sync.WaitGroup,mu *sync.Mutex, counter *int) {
    defer wg.Done()

    for i := 0; i < 100000; i++ {
		mu.Lock()
        (*counter)++
		mu.Unlock()
    }
}

func main() {
    counter := 0

    var wg sync.WaitGroup
	var mu sync.Mutex

    wg.Add(5)

    for i := 0; i < 5; i++ {
        go increment(&wg,&mu, &counter)
    }

    wg.Wait()

    fmt.Println("Final Counter:", counter)
}