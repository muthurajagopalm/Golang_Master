package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "go is a super language known for its concurrency and this concurrency is used to develop lots of tasks very quickly with minimum thread usage"
	words := strings.Fields(sentence)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
