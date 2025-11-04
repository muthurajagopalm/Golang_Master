package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	count := 0
	for val := range ticker.C {
		fmt.Println("tick", val)
		count++

		if count == 5 {
			break
		}

	}
}
