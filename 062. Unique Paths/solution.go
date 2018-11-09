package main

import "fmt"

func uniquePaths(m int, n int) int {
	ways := make([]int, n)
	ways[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			ways[j] += ways[j-1]
		}
	}
	return ways[n-1]
}

// 递归实现
//func findPaths(x int, y int, m int, n int) int {
//	if x == m || y == n {
//		return 1
//	}
//
//	right := findPaths(x + 1, y, m, n)
//	down := findPaths(x, y + 1, m, n)
//	return right + down
//}

func main() {
	res := uniquePaths(3, 2)
	fmt.Println(res)
}
