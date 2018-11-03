package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else {
		ans := make([]int, n)
		ans[0] = 1
		ans[1] = 2
		for i := 2 ; i < n; i++ {
			ans[i] = ans[i-1] + ans[i-2]
		}
		return ans[n-1]
	}
}

func main() {
	n := 2
	res := climbStairs(n)
	fmt.Println(res)
}
