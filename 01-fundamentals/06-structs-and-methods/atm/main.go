package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       int
}

func (acc BankAccount) DisplayAccountInfo() {
	fmt.Println("Account Number:", acc.AccountNumber)
	fmt.Println("Holder Name:", acc.HolderName)
	fmt.Println("Balance:", acc.Balance)
}

func (acc *BankAccount) Deposit(amount int) {
	if amount <= 0 {
		fmt.Println("Amount should be greater than 0")
	} else {
		acc.Balance += amount
		fmt.Println("Amount of", amount, " is sucessfully deposited")
		fmt.Println("Your current Balance is:", acc.Balance)
	}
}

func (acc *BankAccount) Withdraw(amount int) {
	if amount <= 0 {
		fmt.Println("Amount should be greater than 0")
	} else if amount > acc.Balance {
		fmt.Println("Insufficient balance")
	} else {
		acc.Balance -= amount
		fmt.Println("Amount of", amount, " is sucessfully withdrawn")
		fmt.Println("Your current Balance is:", acc.Balance)
	}
}

func main() {

	account := BankAccount{
		AccountNumber: "1234accno9009",
		HolderName:    "Jaison",
		Balance:       100000,
	}

	for {

		fmt.Println("---Menu---")
		fmt.Println("1. Display Account info")
		fmt.Println("2. Deposit Amount")
		fmt.Println("3. Withdraw Amount")
		fmt.Println("4. Exit")
		fmt.Println("----------")

		var choice int

		fmt.Print("Enter your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			account.DisplayAccountInfo()
		case 2:
			var amount int
			fmt.Print("Enter amount to deposit:")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 3:
			var amount int
			fmt.Print("Enter amount to withdraw:")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 4:
			fmt.Println("Thanks for Using our GO ATM")
			return
		default:
			fmt.Println("Invalid input")
			continue
		}
	}
}
