package main

import "fmt"

// struct
type Student struct {
	ID         int
	Name       string
	Email      string
	Department string
}

// method
func (s Student) Display() {
	fmt.Println("Student Name:", s.Name)
	fmt.Println("Student Department:", s.Department)
}

func (s *Student) UpdateName(name string) { // only (s Student) it will copy of the struct. Changes to s won't modify the original student so use pointer receiver
	s.Name = name
}

func main() {

	student := Student{
		ID:         1,
		Name:       "Jaison",
		Email:      "connect@jaisondavid.in",
		Department: "CSE",
	}

	fmt.Println(student)
	fmt.Println(student.Name)
	student.Name = "David"
	fmt.Println(student.Name)

	student.Display() // method calling

	student.UpdateName("new name")
	student.Display()

}

//default values if data is not given
// int 		- 	0
// string	- 	""
// bool		-	false
// Pointer	-	nil
