package main

import "fmt"

func main() {
	balance := 5000

	for {
		fmt.Println("\n===== GO ATM =====")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: ₹%d\n", balance)

		case 2:
			var amount int

			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount. Deposit must be greater than zero.")
			} else {
				balance += amount
				fmt.Printf("₹%d deposited successfully.\n", amount)
				fmt.Printf("Updated balance: ₹%d\n", balance)
			}

		case 3:
			var amount int

			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount. Withdrawal must be greater than zero.")
			} else if amount > balance {
				fmt.Println("Insufficient balance.")
			} else {
				balance -= amount
				fmt.Printf("₹%d withdrawn successfully.\n", amount)
				fmt.Printf("Remaining balance: ₹%d\n", balance)
			}

		case 4:
			fmt.Println("Thank you for using Go ATM!")
			return

		default:
			fmt.Println("Invalid choice. Please select 1-4.")
		}
	}
}
