package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	lenNums := len(nums)
	if lenNums < 3 {
		return res
	}
	l, r, dif := 0, 0, 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		// 避免重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l = i + 1
		r = lenNums - 1
		dif = -nums[i]
		for l < r {
			if nums[l]+nums[r] == dif {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 避免重复和超出索引
				for (nums[l] == nums[l-1]) && l < r {
					l++
				}
				for (nums[r] == nums[r+1]) && r > 0 {
					r--
				}
			} else if nums[l]+nums[r] < dif {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
