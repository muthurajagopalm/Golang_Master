package main

import "fmt"

func main() {
	s := "Hello Muthu!, you are doing great"
	count := 0

	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	fmt.Println(count)
}
