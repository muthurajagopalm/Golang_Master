package main

import (
	"fmt"
	"sync"
)

func Workers(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for nums := range jobs {
		results <- nums * nums
	}
}

func main() {

	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Workers(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)

	}()

	for res := range results {
		fmt.Println("The received values from the channel are", res)
	}
}
