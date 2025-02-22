package main

import (
	"time"
)

/*Go no permite la herencia, go utiliza la composicion. la composicion,
a diferencia de la herencia, no es una clase hija de... sino que
contiene los metodos de las clases indicadas.*/

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	//select * from persona where.
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	//select * from emplyee where.
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEployee, error) {
	var ftEmployee FullTimeEployee
	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p
	return ftEmployee, nil

}
