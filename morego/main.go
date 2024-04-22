package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 4, 23, 15, 500}

	maxValue := nums[0]

	for _, n := range nums[1:] {
		if n > maxValue {
			maxValue = n
		}
	}
	fmt.Println(maxValue)

}
