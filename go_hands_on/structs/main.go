package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Muthu", Age: 25},
		{Name: "Naveen", Age: 27},
		{Name: "DhanyaShree", Age: 5},
		{Name: "Aathirai", Age: 7},
		{Name: "Krithik", Age: 6},
		{Name: "Yugan", Age: 4},
		{Name: "Lakshika", Age: 1},
	}

	for _, person := range people {
		if person.Age > 18 {
			fmt.Println(person.Name)
		}

	}
}
