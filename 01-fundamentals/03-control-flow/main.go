package main

import "fmt"

func main() {

	age := 19

	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not Eligible to vote")
	}

	status := "SUCCESS"

	switch status {
	case "PENDING", "PROCESSSING":
		fmt.Println("Order is being processing")
	case "SUCCESS":
		fmt.Println("Order completed")
	case "FAILED":
		fmt.Println("Order failed")
	default:
		fmt.Println("Unkown status")
	}

	marks := 85

	switch {
	case marks >= 90:
		fmt.Println("Grade A")

	case marks >= 75:
		fmt.Println("Grade B")

	default:
		fmt.Println("Grade C or below")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 { // go doesn't have keyword while
		fmt.Println(i)
		i++
	}

	// for {
	// 	fmt.Println("Infinite loop running...")
	// }

	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

}
