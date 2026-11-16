package main

import "fmt"

func main() {

	userCh := make(chan string, 1)
	paymentCh := make(chan string, 1)

	go func() {
		userCh <- "User Created Successfully"
	}()

	go func() {
		paymentCh <- "Payment completed"
	}()

	select {
	case user := <-userCh:
		fmt.Println(user)
	case payment := <-paymentCh:
		fmt.Println(payment)
	}

}