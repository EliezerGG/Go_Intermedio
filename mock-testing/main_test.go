package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  30,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
				Person: Person{
					Name: "John Doe",
					Age:  30,
					DNI:  "1",
				},
			},
		},
	}

	for _, tt := range table {
		employee, err := GetFullTimeEmployeeById(tt.id, tt.dni)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if employee != tt.expectedEmployee {
			t.Errorf("expected %v, got %v", tt.expectedEmployee, employee)
		}
	}
}
