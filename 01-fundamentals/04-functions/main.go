package main

import "fmt"

func greet(name string) { // In Go, lowercase functions are accessible only within the same package, while uppercase functions are exported and can be accessed from other packages through imports.
	fmt.Println("Welcome!", name)
}

func add(a int, b int) int { //can also add(a,b int)
	return a + b
}

func calculate(a, b int) (int, int) { //Go natively supports multiple return values from functions, while languages like Python, JavaScript, Rust, and Swift use tuples, arrays, or other structures to achieve it.
	sum := a + b
	diff := a - b
	return sum, diff
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}

func calc(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	greet("jaison")

	result := add(12, 90)
	fmt.Println(result)

	sum, difference := calculate(20, 10)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)

	DivResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", DivResult)

}
