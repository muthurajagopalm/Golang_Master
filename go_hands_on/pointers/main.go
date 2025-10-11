package main

import "fmt"

func sumNums(a, b *int) {
	*a, *b = *b, *a

}

func main() {

	x, y := 10, 20
	fmt.Println("Before Swap:", x, y)

	sumNums(&x, &y)
	fmt.Println("After Swap:", x, y)

}
