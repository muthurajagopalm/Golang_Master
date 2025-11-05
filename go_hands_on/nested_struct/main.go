package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Employee struct {
	ID      int
	Name    string
	Address Address
}

func main() {
	emp := Employee{
		ID:   1,
		Name: "Muthu",
		Address: Address{
			City:  "Chennai",
			State: "TamilNadu",
		},
	}

	fmt.Println("Employee ID;", emp.ID)
	fmt.Println("Employee Name;", emp.Name)
	fmt.Println("Employee City;", emp.Address.City)
	fmt.Println("Employee State;", emp.Address.State)
}
