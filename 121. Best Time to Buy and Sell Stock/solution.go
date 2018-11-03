package main

import "fmt"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 1 {
		return 0
	}
	min := prices[0]
	max := 0
	for _, value := range prices {
		if min > value {
			min = value
		} else {
			if max < value - min {
				max = value - min
			}
		}
	}
	return max
}

func main() {
	res := maxProfit([]int{7,1,5,3,6,4})
	fmt.Println(res)
}
