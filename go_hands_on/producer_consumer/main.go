package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("The producer %d\n", i)
		ch <- i * i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Printf("The consumer received %d\n", val)
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
	fmt.Println("Successfully Done!")
}
