package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("eg.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
