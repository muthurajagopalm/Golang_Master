package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	var cleaned []rune

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned = append(cleaned, unicode.ToLower(ch))
		}
	}

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	result := IsPalindrome(s)

	if result {
		fmt.Println("The given sentence is a palindrome")
	} else {
		fmt.Println("The given sentence is not a palindrome")
	}
}
