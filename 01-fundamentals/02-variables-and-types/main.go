package main

import "fmt"

func main() {

	const appName = "AppName fixed" // it cannot be later modified anywhere

	const (
		statusPending = "PENDING"
		statusSuccess = "SUCCESS"
		statusFailed  = "FAILED"
	)

	var name string = "jaison"
	var age int = 18
	var salary float64 = 25000.00

	var name2 = "jaison" // auto assign the type of the variable based on the give value
	var age2 = 18

	name3 := "jaison" // only can declared inside the functions . local only

	_ = age // to prevent declared and not used error

	var name_empty string //auto assing empty . here ""
	var is_true bool      // here default false like for all not assigned variables

	_ = name2
	_ = age2
	_ = name3
	_ = name_empty
	_ = is_true

	fmt.Println(name)
	fmt.Println(salary)

}
