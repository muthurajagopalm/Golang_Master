package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("eg.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	fmt.Println("Read the file successfully", string(file))
}
