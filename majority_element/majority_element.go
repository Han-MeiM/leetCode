package main

import "fmt"

func majorityElement(nums []int) int {
	count, max, length := make(map[int]int), 0, len(nums)
	for _, value := range nums {
		count[value] = count[value] + 1
		if count[value] > length/2 {
			max = value
		}
	}
	return max
}

func main() {
	nums := []int{6,5,5}
	res := majorityElement(nums)
	fmt.Println(res)
}
