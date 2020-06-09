package main

import "fmt"

func containsDuplicate(nums []int) bool {
	container := make(map[int]int)
	for _, v := range nums {
		if container[v] > 0 {
			return true
		} else {
			container[v] = 1
		}
	}
	return false
}

func main() {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	res := containsDuplicate(nums)
	fmt.Println(res)
}
