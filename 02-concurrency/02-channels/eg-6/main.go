package main

import "fmt"

func main() {

	emailCh := make(chan string, 1)
	smsCh := make(chan string, 1)

	go func() {
		emailCh <- "Email sent"
	}()

	go func() {
		smsCh <- "SMS sent"
	}()

	select {
	case email := <-emailCh:
		fmt.Println(email)
	case sms := <-smsCh:
		fmt.Println(sms)
	default:
		fmt.Println("No message available")
	}

}