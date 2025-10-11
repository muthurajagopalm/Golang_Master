package main

import (
	"fmt"
	"sync"
)

func sendValue(val int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- val
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(3)

	go sendValue(10, ch, &wg)
	go sendValue(20, ch, &wg)
	go sendValue(30, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for values := range ch {
		fmt.Println("Received:", values)
	}

}
