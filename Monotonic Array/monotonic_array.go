package main

import (
	"fmt"
)

func isMonotonic(A []int) bool {
	length := len(A)
	if length == 1 {
		return true
	}

	// 1 为递增，-1 为递减
	previous := 0
	for i := 1; i < length; i++ {
		curren := 0

		if A[i] > A[i - 1] {
			curren = 1
		} else if A[i] == A[i - 1] {
			continue
		} else {
			curren = -1
		}

		if previous == 0 {
			previous = curren
		} else if previous != curren {
			return false
		}
	}

	return true
}

func main() {
	nums := []int{1,2,1}
	fmt.Println(isMonotonic(nums))
}
