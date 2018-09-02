package main

import "fmt"

func removeElement(nums []int, val int) int {
	m := 0
	for _, n := range nums {
		if n != val {
			nums[m] = n
			m++
		}
		fmt.Println(nums)
	}
	return m
}

func main() {
	nums, val := []int{3,2,2,3}, 3
	res := removeElement(nums, val)
	fmt.Println(nums, res)
}
