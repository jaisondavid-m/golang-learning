
package main

import (
    "fmt"
    "sync"
)

func deposit(mu *sync.Mutex, wg *sync.WaitGroup, balance *int, amount int) {
    defer wg.Done()

    // mu.Lock()
    *balance += amount
    // mu.Unlock()
}

func main() {
    balance := 1000

    var mu sync.Mutex
    var wg sync.WaitGroup

    wg.Add(5)

    go deposit(&mu, &wg, &balance, 1000)
    go deposit(&mu, &wg, &balance, 1000)
	go deposit(&mu, &wg, &balance, 1000)
	go deposit(&mu, &wg, &balance, 1000)
	go deposit(&mu, &wg, &balance, 1000)

    wg.Wait()

    fmt.Println("Final balance:", balance)
}
