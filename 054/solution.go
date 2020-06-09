package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	yLen := len(matrix)
	if yLen == 0 {
		return []int{}
	}
	if yLen == 1 {
		return matrix[0]
	}

	xLen, key := len(matrix[0]), 0
	sum := yLen * xLen
	res := make([]int, sum)
	yLen, xLen = yLen-1, xLen-1

	x, y, circle := 0, 0, 0
	for key < sum {
		res[key] = matrix[y][x]
		key++
		if y == circle && x != xLen-circle {
			x++
			continue
		} else if x == xLen-circle && y != yLen-circle {
			y++
			continue
		} else if y == yLen-circle && x != circle {
			x--
			continue
		} else if x == circle && y != circle+1 {
			y--
			continue
		}
		x++
		circle++
	}
	return res
}

func main() {
	item1 := []int{1, 2, 3, 4}
	item2 := []int{5, 6, 7, 8}
	item3 := []int{9, 10, 11, 12}
	arr := [][]int{item1, item2, item3}
	res := spiralOrder(arr)
	fmt.Println(res)
}
