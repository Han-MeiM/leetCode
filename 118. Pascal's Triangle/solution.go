package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)

		if i == 0 {
			row[0] = 1
		} else if i == 1 {
			row[0] = 1
			row[1] = 1
		} else {
			for j := range row {
				if j == 0 || j == len(row)-1 {
					row[j] = 1
					continue
				}
				row[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res[i] = row
	}
	return res
}

func main() {
	res := generate(4)
	fmt.Println(res)
}
