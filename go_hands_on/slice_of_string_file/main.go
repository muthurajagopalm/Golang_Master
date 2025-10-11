package main

import (
	"fmt"
	"os"
)

func main() {

	words := []string{"Go is awesome", "Interfaces are powerful", "File handling is easy"}

	file, err := os.Create("eg.txt")
	if err != nil {
		fmt.Println("Error creating the file", err)
		return
	}
	fmt.Println("File Created successfully", file)
	defer file.Close()

	for _, word := range words {
		_, err := file.WriteString(word + "\n")

		if err != nil {
			fmt.Println("Error writing the words", err)
			return
		}
	}
	fmt.Println("Slice written to the file successfully")

}
