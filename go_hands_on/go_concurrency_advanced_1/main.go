package main

import (
	"fmt"
	"sync"
)

func worker(id, start, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Printf("The values that are working are %d\n:", i)
		ch <- i * i
	}

}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go worker(1, 1, 5, ch, &wg)

	wg.Add(1)
	go worker(2, 6, 10, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println("The values received from the channel are:", val)
	}

	fmt.Println("Everything is done")
}
