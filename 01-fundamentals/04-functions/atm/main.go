package main

import "fmt"

func display_menu() {
	fmt.Println("\n===== GO ATM =====")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}

func get_balance(amount int) {
	fmt.Println("Your current balance is:", amount)
}

func deposit_amount(total, amount int) (int, error) {
	if amount <= 0 {
		return total, fmt.Errorf("the amount should be greater than 0")
	}

	total += amount
	fmt.Println("Amount of", amount, "is deposited successfully")
	get_balance(total)
	return total, nil
}

func withdraw_amount(total, amount int) (int, error) {
	if amount <= 0 {
		return total, fmt.Errorf("the amount should be greater than 0")
	} else if amount > total {
		return total, fmt.Errorf("insufficient balance")
	}

	total -= amount
	fmt.Println("Amount of", amount, "is withdrawn successfully")
	get_balance(total)
	return total, nil
}

func start_greet() {
	fmt.Println("Welcome to GO ATM")
}

func end_greet() {
	fmt.Println("Thank you for using Go ATM")
}

func main() {
	balance := 50000
	start_greet()

	for {
		display_menu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			get_balance(balance)

		case 2:
			var amount int
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&amount)

			var err error
			balance, err = deposit_amount(balance, amount)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			var amount int
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&amount)

			var err error
			balance, err = withdraw_amount(balance, amount)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			end_greet()
			return

		default:
			fmt.Println("Invalid choice. Please select 1-4.")
		}
	}
}
