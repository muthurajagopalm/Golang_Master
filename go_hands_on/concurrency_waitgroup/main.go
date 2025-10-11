package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Microsecond)
	}
}

func PrintLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println("Letters:", string(i))
		time.Sleep(500 * time.Microsecond)
	}
}

func printName(wg *sync.WaitGroup) {
	defer wg.Done()
	for n := 1; n <= 5; n++ {
		fmt.Println(n, "Muthu Perumal Rajagopal")
		time.Sleep(500 * time.Microsecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go printNumbers(&wg)
	go PrintLetters(&wg)
	go printName(&wg)

	wg.Wait()
}
