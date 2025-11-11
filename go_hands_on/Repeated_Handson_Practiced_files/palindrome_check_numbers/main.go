package main

import "fmt"

func IsPalindrome(num int) bool {
	original := num
	reversed := 0

	for num > 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num = num / 10
	}
	return original == reversed
}

func main() {
	n := 12321
	result := IsPalindrome(n)

	if result {
		fmt.Println("The given number is palindrome")
	} else {
		fmt.Println("The given number is not a palindrome")
	}
}
