package main

import "fmt"

func main() {
	sentence := "golang is a good language"

	charCount := make(map[rune]int)

	for _, charac := range sentence {
		charCount[charac]++

	}

	for charac, count := range charCount {
		fmt.Printf("%q: %d\n", charac, count)
	}
}
