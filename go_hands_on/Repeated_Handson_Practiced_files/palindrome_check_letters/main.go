package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {

	str := strings.ToLower(s)

	left := 0
	right := len(s) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "RaceCare"
	result := IsPalindrome(s)
	if result {
		fmt.Println("The given string is a palindrome")
	} else {
		fmt.Println("THe given string is not a palindrome")
	}
}
