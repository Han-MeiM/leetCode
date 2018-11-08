package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	right := 1

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1]
		result[i] *= right
	}

	return result
}

func main() {
	nums := []int{1,2,3,4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
