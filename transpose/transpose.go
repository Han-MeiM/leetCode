package main

import "fmt"

func transpose(A [][]int) [][]int {
	y_len := len(A)
	x_len := len(A[0])
	res := make([][]int, x_len)
	for n := 0; n < x_len; n++ {
		row := make([]int, y_len)
		for m := 0; m < y_len; m++ {
			row[m] = A[m][n]
		}
		res[n] = row
	}
	return res
}

func main() {
	A := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	res := transpose(A)
	fmt.Println(res)
}
