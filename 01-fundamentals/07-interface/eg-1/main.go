package main

import "fmt"

type UPI struct{}

func (u UPI) Pay(amount int) {
	fmt.Println("Paid using UPI:",amount)
}

type CreditCard struct{}

func (c CreditCard) Pay(amount int) {
	fmt.Println("Paid using Credit Card:", amount)
}

type PaymentMethod interface {
	Pay(amount int)
}

func processPayment(payment PaymentMethod, amount int) {
	payment.Pay(amount)
}

func main() {
	
	upi := UPI{}
	card := CreditCard{}

	processPayment(upi,500)
	processPayment(card,1200)

}