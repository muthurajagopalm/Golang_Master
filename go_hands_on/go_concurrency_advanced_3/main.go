package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	ch <- 10
	ch <- 20

	fmt.Println("Done successfully")
}
