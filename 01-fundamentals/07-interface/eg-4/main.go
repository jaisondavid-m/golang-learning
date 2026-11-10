package main

import "fmt"

func printValue(value any) { // here any is empty interface = interface{}

	fmt.Println(value)

	number, ok := value.(int)
	if ok {
		fmt.Println(number + 10)
	} else {
		fmt.Println("Given value is not integer")
	}

}

func checkValue(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	case float64:
    	fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {

	printValue("Go Backend")
	printValue(25)
	printValue(true)
	printValue(12.89)

	checkValue(25)
	checkValue("Go Backend")
	checkValue(true)
	checkValue(12.89)
}
