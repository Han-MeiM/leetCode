package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, column := len(nums), len(nums[0])

	if row * column != r * c {
		return nums
	}

	res, res_row, cindex := [][]int{}, []int{}, 0

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			res_row = append(res_row, nums[i][j])
			if cindex++; cindex == c {
				res = append(res, res_row)
				res_row = []int{}
				cindex = 0
			}
		}
	}
	
	return res
}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	r, c := 2, 2
	res := matrixReshape(nums, r, c)
	fmt.Println(res)
}
