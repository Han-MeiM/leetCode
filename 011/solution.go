package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		area := 0
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left = left + 1
		} else {
			area = (right - left) * height[right]
			right = right - 1
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(arr)
	fmt.Println(res)
}
