package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		3: "Orange",
		2: "Pomo",
		1: "Apple",
		4: "kiwi",
		5: "Amla",
		8: "Dragon fruit",
		9: "Banana",
	}

	keys_slice := make([]int, 0, len(m))

	for k := range m {
		keys_slice = append(keys_slice, k)
	}
	sort.Ints(keys_slice)
	fmt.Println(keys_slice)

	for _, val := range keys_slice {
		fmt.Println(val, ":", m[val])
	}
}
