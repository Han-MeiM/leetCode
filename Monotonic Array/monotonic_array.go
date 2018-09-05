package main

import (
	"fmt"
)

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	return true
}

func main() {
	nums := []int{3,2,2,1}
	fmt.Println(isMonotonic(nums))
}
