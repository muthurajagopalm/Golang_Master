package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 90
		ch <- 20
		close(ch)

	}()
	for val := range ch {
		fmt.Println("Received ", val)

	}
}
