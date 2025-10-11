package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("eg.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nHello! ...etc.,")
	if err != nil {
		fmt.Println("Error appending the file", err)
		return
	}

	fmt.Println("Appended the file successfully", file)

	data, err := os.ReadFile("eg.txt")
	if err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	fmt.Println("Content of the file is displayed below\n", string(data))

}
