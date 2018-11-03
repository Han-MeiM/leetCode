package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxL, tempL := 0, 0
	for _, v := range nums {
		if v == 1 {
			tempL ++
			if maxL < tempL {
				maxL = tempL
			}
		} else {
			tempL = 0
		}
	}

	return maxL
}

func main() {
	nums := []int{1,1,0,1,1,1}
	res := findMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
