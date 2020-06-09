package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}
