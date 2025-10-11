package main

import "fmt"

func main() {
	arr := []int{1, 3, 6, 9, 1, 18}
	sum := 0

	for _, val := range arr {
		sum = sum + val
	}
	fmt.Println(sum)
}
