package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	sort_nums := sort.IntSlice(nums)
	sort_nums.Sort()
	res := search(sort_nums, 0, []int{})
	return res
}

func search(nums []int, index int, temp []int) [][]int {
	ret := make([][]int, 0)

	if index == len(nums) {
		return append(ret, temp)
	}

	leftTemp := make([]int, len(temp))
	copy(leftTemp, temp)
	leftArr := search(nums, index + 1, append(leftTemp, nums[index]))

	rightTemp := make([]int, len(temp))
	copy(rightTemp, temp)
	rightArr := search(nums, index + 1, rightTemp)

	ret = append(ret, leftArr...)
	ret = append(ret, rightArr...)
	return  ret
}

//func subsets(nums []int) [][]int {
//	ret := make([][]int, 0)
//	ret = append(ret, make([]int, 0))
//	if len(nums) == 0 {
//		return ret
//	}
//
//	if len(nums) == 1 {
//		return append(ret, []int{nums[0]})
//	}
//	for k, v := range nums {
//		subset := subsets(nums[:k])
//		for _, item := range subset {
//			tmp := make([]int, len(item))
//			tmp = item
//			tmp = append(tmp, v)
//			ret = append(ret, tmp)
//		}
//	}
//
//	return ret
//}

func main() {
	nums := []int{0,3,5,7,9}
	res := subsets(nums)
	fmt.Println(res)
}
