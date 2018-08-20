package main

import (
	"sort"
	"fmt"
)

func arrayPairSum(nums []int) int {
	res := 0
	// 先将数组进行排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func main() {
	nums := []int{3, 4, 6, 10, 1, 4, 14, 12, 10}
	res := arrayPairSum(nums)
	fmt.Println(res)
}
