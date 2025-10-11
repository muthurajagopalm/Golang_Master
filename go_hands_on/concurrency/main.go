package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Numbers", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println("Letters", string(ch))
		time.Sleep(500 * time.Millisecond)
	}
}

func printMyName() {
	for n := 1; n <= 5; n++ {
		fmt.Println(n, "Muthu Perumal Rajagopal")
		time.Sleep(500 * time.Millisecond)

	}
}
func main() {

	go printNumbers()
	go printLetters()
	go printMyName()

	time.Sleep(3 * time.Second)
	fmt.Println("Main execution finishedd")

}
