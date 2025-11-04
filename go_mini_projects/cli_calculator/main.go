package main

import "fmt"

func add(a, b float64) float64 {
	return a + b

}
func subtract(a, b float64) float64 {
	return a - b

}

func multiply(a, b float64) float64 {
	return a * b

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil

}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the Operator + - * /")
	fmt.Scanln(&operator)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	var result float64
	var err error

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result, err = divide(num1, num2)
	default:
		fmt.Println("Invalid Operators")
	}

	if err != nil {
		fmt.Println("Error", err)

	} else {
		fmt.Println("Result:", result)
	}
}
