package main

import (
	"fmt"
	// "errors"
)

func withdraw(balance,amount int) (int, error) {
	if amount <= 0 {
		return 0,fmt.Errorf("The amount should be greater than 0")
	}
	if amount > balance {
		return 0,fmt.Errorf(
			"Insufficient balance: requested %d, available %d",
			amount,balance,
		)
	}
	return balance - amount, nil
}

func display(balance int, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The updated balance is:",balance)
	}
}

func main() {
	display(withdraw(5000, 1000)) // Success: 4000
	display(withdraw(5000, 6000)) // Insufficient balance
	display(withdraw(5000, -100)) // Invalid amount
}