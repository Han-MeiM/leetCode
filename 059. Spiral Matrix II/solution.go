package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := makeArr(n)
	square := n*n
	num := 1
	var x, y int
	for num <= square {
		// 右
		for y < n && res[x][y] == 0 {
			res[x][y] = num
			y++
			num++
		}
		y--
		x++
		// 下
		for x < n && res[x][y] == 0 {
			res[x][y] = num
			x++
			num++
		}
		x--
		y--
		// 左
		for y >= 0 && res[x][y] == 0 {
			res[x][y] = num
			y--
			num++
		}
		x--
		y++
		// 上
		for x >= 0 && res[x][y] == 0 {
			res[x][y] = num
			x--
			num++
		}
		y++
		x++
	}
	return res
}

func makeArr(n int) [][]int {
	arr1 := make([][]int, n)
	for index := range arr1 {
		arr2 := make([]int, n)
		arr1[index] = arr2
	}
	return arr1
}

func main() {
	n := 3
	res := generateMatrix(n)
	fmt.Println(res)
}
