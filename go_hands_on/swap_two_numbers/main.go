package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	x, y := 10, 20
	fmt.Println("Before Swap:", x, y)

	x, y = swap(x, y)
	fmt.Println("After Swap:", x, y)
}
