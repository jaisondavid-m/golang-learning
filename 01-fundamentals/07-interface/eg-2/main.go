package main

import "fmt"

type Bird struct{}
type Human struct{}
type Car struct{}

func (b Bird) Move() {
	fmt.Println("Bird is Flying")
}

func (r Human) Move() {
	fmt.Println("Human is Walking")
}

func (c Car) Move() {
	fmt.Println("Car is Moving")
}

type Mover interface{
	Move()
}

func StartMoving(m Mover) {
	m.Move()
}

func main() {

	bird := Bird{}
	human := Human{}
	car := Car{}

	StartMoving(bird)
	StartMoving(human)
	StartMoving(car)

}