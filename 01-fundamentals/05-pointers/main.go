package main

import "fmt"

func updateAge(age *int) {
	*age = 30
}

func updateBalance(balance *int, amount int) {
	*balance += amount
}

func main() {

	age := 19

	// var ptr *int = &age
	ptr := &age

	*ptr = 25

	updateAge(&age)

	fmt.Println("Value:", age)
	fmt.Println("Address:", &age)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Dereferenced value:", *ptr)

	var ptr1 *int
	fmt.Println(ptr1)

	if ptr1 != nil {
		fmt.Println(*ptr1)
	} else {
		fmt.Println("Pointer is nil")
	}

	balance := 10000
	updateBalance(&balance, 1000)
	updateBalance(&balance, 4000)
	fmt.Println("Final balance:", balance)

}
