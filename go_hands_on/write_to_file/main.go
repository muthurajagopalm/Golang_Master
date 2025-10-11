package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("eg.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return

	}
	defer file.Close()

	_, err = file.WriteString("\n Hey Muthu!, You are Blessed Beyond Measure!")
	if err != nil {
		fmt.Println("Error Writing to the file", err)
		return
	}
	fmt.Println("File Created Sucessfully", file)

}
