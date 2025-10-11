package main

import "fmt"

func smallNums(arr []int) int {
	small := arr[0]

	for _, val := range arr {
		if val < small {
			small = val
		}
	}
	return small
}

func largeNums(arr []int) int {
	large := arr[0]

	for _, val := range arr {
		if val > large {
			large = val
		}
	}
	return large
}

func main() {
	arr := []int{2, 5, 87, 99, 18, 10000}

	identSmall := smallNums(arr)
	fmt.Println(identSmall)

	identLarge := largeNums(arr)
	fmt.Println(identLarge)

}
