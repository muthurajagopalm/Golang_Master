package main

import "fmt"

func main() {

	a := 10
	b := &a
	*b = 99

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)

}
