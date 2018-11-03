package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for k := range nums {
		ret := []int{}
		ret = append(ret, nums[k])
		temp := handlePermute(append(nums[:k:k], nums[k+1:]...), ret)
		res = append(res, temp...)
	}
	return res
}

func handlePermute(nums []int, ret []int) [][]int {
	res := [][]int{}
	if len(nums) == 1 {
		ret = append(ret, nums[0])
		return [][]int{ret}
	}
	for k := range nums {
		retTemp := append(ret, nums[k])

		// 组合剩余的数字为数组
		item := make([]int, len(nums[:k]))
		copy(item, nums[:k])
		item = append(item, nums[k+1:]...)

		temp := handlePermute(item, retTemp)
		res = append(res, temp...)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := permute(nums)
	fmt.Println(res)
}
