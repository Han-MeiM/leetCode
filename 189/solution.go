package main

import "fmt"

func reverse(nums []int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp

		left++
		right--
	}
}

// 1. 整个反转一次
// 2. 再根据 k 作为分界线，两边各自反转一次
func rotate(nums []int, k int) {
	// 避免 k 超过 nums 的长度
	k %= len(nums)

	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])

	fmt.Println(nums)
}

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums := []int{-1}
	rotate(nums, 2)
}
