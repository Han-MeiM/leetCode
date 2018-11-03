package main

import "fmt"

//func moveZeroes(nums []int) []int {
//	length := len(nums)
//	for k := 0; k < length; k++ {
//		for k_r := k + 1 ; k_r < length; k_r++ {
//			if nums[k] == 0 {
//				nums[k] = nums[k_r]
//				nums[k_r] = 0
//				continue
//			}
//
//			if nums[k] > nums[k_r] {
//				if nums[k_r] == 0 {
//					continue
//				}
//				temp := nums[k]
//				nums[k] = nums[k_r]
//				nums[k_r] = temp
//			}
//		}
//	}
//	return nums
//}

func moveZeroes(nums []int) []int {
	length := len(nums)
	for k := 0; k < length; k++ {
		for k_r := k + 1 ; k_r < length; k_r++ {
			if nums[k] == 0 {
				nums[k] = nums[k_r]
				nums[k_r] = 0
				continue
			}
		}
	}
	return nums
}

func main() {
	nums := []int{0,1,0,3,12}
	res := moveZeroes(nums)
	fmt.Println(res)
}
