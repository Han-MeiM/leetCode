package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	left, right, mid := 0, len(nums)-1, 0

	for left < right {
		mid = (left + right) / 2
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if nums[mid] < target {
				left = mid + 1
			} else {
				if target > nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else {
			if target > nums[mid] {
				if target < nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	arr := []int{0,1,2,4,5,6,7}
	res := search(arr, 4)
	fmt.Println(res)
}
