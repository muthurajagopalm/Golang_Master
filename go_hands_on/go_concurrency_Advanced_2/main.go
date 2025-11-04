package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 4
	ch <- 2

	fmt.Println("Printing the value:", <-ch)
	fmt.Println("Printing the value:", <-ch)
}
