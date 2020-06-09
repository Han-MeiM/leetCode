package main

import "fmt"

func maxProfit(prices []int) int {
	ans := 0

	if len(prices) <= 1 {
		return 0
	}

	for k := range prices {
		if k == 0 {
			continue
		}
		if prices[k] - prices[k-1] >= 0 {
			ans += prices[k] - prices[k-1]
		}
	}
	return ans
}

func main() {
	arr := []int{7,1,5,3,6,4}
	res := maxProfit(arr)
	fmt.Println(res)
}
