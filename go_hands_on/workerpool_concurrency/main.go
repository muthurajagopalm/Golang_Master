package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		fmt.Printf("Worker %d Processing %d\n", id, num)
		result <- num * num
	}
}

func main() {
	var wg sync.WaitGroup
	numJobs := 5
	numWorker := 3
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= numWorker; w++ {
		wg.Add(1)
		go worker(w, jobs, result, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println("Received", res)
	}
}
