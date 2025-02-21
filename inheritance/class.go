package main

import "fmt"

type Person struct {
	name string
	age  int
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "TemporaryEmployee"
}

type PrintInfo interface {
	getMessage() string
}
type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (FullTimeEmployee FullTimeEmployee) getMessage() string {
	return "FullTimeEmployee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	fmt.Printf("%+v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
