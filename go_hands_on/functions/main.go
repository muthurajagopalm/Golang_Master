package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return a / b, nil
}

func main() {

	result, err := Divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = Divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
