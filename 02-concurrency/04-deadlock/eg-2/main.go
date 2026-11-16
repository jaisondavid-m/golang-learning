package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex

	mu.Lock()
	fmt.Println("First Lock")
	mu.Lock()
	fmt.Println("Second Lock") // error because already it lock a value to hold another first it need unlock the first one
	mu.Unlock()

}