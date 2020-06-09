package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	result := 0

	if num < 10 {
		return 1
	}

	if num < 26 {
		return 2
	}

	if num < 100 {
		return 1
	}

	numStr := strconv.Itoa(num)

	for i := 1; i < 3; i++ {
		left, _ := strconv.Atoi(numStr[:i])
		right, _ := strconv.Atoi(numStr[i:])

		if left < 26 {
			result = result + translateNum(right)
		}
	}

	return result
}

func main() {
	result := translateNum(12258)

	fmt.Println(result)
}
