package main

import "fmt"

func missingNumber(nums []int) int {
	sum := len(nums)
	for i, v := range nums {
		sum = sum + i - v
	}
	return sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
