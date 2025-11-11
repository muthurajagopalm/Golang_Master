package main

import "fmt"

func RemoveElements(nums []int, val int) int {
	k := 0

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == val {
			nums[k] = nums[i]
			k++
		}

	}
	fmt.Println(k)
	return k

}

func main() {

	nums := []int{3, 2, 2, 3}
	val := 3
	RemoveElements(nums, val)

}
