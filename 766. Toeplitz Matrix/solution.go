package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	x, y := len(matrix[0]), len(matrix)

	if x == 1 || y == 1 {
		return true
	}

	// 默认有两行两列以上
	for i := 0; i < x + y -1; i++ {
		if i == 0{
			continue
		}
		
		topX, topY := 0, y - 1 - i
		btmX, btmY := i, y - 1
		
		if i > x - 1 {
			btmX = x - 1
			btmY = y - 1 - (i - (x - 1))
		}

		if i > y - 1 {
			topX = i - (y - 1)
			topY = 0
		}

		if matrix[btmY][btmX] != matrix[topY][topX] {
			return false
		}

		for j := 1; j <= btmX - topX; j++ {
			if matrix[btmY][btmX] != matrix[topY + j][topX + j] {
				return false
			}
		}
	}

	return true
}

func main() {
	matrix := [][]int{
		{76,17,97,63,74,73,47,55},
		{90,76,17,97,63,74,73,47},
		{17,90,76,17,97,63,74,73},
		{18,17,90,76,17,97,63,74},
		{17,18,17,90,76,17,97,0},
		{68,17,18,17,90,76,17,97},
	}
	res := isToeplitzMatrix(matrix)
	fmt.Println(res)
}
