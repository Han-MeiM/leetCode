package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	QuickSort(nums)
	return nums[len(nums)-k]
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	pivot, i := values[0], 1
	left, right := 0, len(values)-1
	for left < right {
		if values[i] > pivot {
			values[i], values[right] = values[right], values[i]
			right--
		} else {
			values[i], values[left] = values[left], values[i]
			left++
			i++
		}
	}
	QuickSort(values[:left])
	QuickSort(values[left+1:])
}

func main() {
	arr := []int{3,1,4,2,6}
	res := findKthLargest(arr, 4)
	fmt.Println(res)
}
