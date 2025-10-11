package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("The number %d is an even\n", i)
		} else {
			fmt.Printf("The number %d is an odd\n", i)
		}
	}
}
