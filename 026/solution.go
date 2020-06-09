package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	m := 1
	for n := 1; n < length; n++ {
		if nums[n] != nums[n-1] {
			nums[m] = nums[n]
			m++
		}
	}

	return m
}

func main() {
	arr := []int{1,1,2}
	res := removeDuplicates(arr)
	fmt.Println(res, arr)
}
