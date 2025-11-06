package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		5: "Orange",
		3: "kiwi",
		8: "Guava",
		1: "Pomogranete",
		2: "Amla",
	}

	type slice_kv struct {
		Key   int
		value string
	}

	var keyval []slice_kv

	for k, v := range m {
		keyval = append(keyval, slice_kv{Key: k, value: v})
		fmt.Println(keyval)
	}

	sort.Slice(keyval, func(i, j int) bool {
		return keyval[i].value < keyval[j].value
	})

	for _, val := range keyval {
		fmt.Println(val.Key, ":", val.value)
	}
}
