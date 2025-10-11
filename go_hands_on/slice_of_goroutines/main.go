package main

import (
	"fmt"
	"sync"
)

func sendValue(val int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := val * val
	ch <- result
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	values := []int{19, 40, 40, 20, 10, 99}
	wg.Add(len(values))

	for _, value := range values {
		go sendValue(value, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Received", v)
	}
}
