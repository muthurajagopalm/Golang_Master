package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("eg.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")

	lineCount := 0
	wordCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	fmt.Println("Number of lines:", lineCount)
	fmt.Println("Number of words:", wordCount)

}
