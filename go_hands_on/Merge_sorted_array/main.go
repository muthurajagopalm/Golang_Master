package main

import "fmt"

func mergeSortArray(num1 []int, m int, num2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if num1[i] > num2[j] {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
		k--
	}

	for j >= 0 {

		num1[k] = num2[j]
		j--
		k--
	}
	fmt.Println(num1)
}

func main() {
	nums_1 := []int{1, 2, 3, 0, 0, 0, 0}
	nums_2 := []int{2, 5, 6, 9}
	m := 3
	n := 4
	mergeSortArray(nums_1, m, nums_2, n)

}
