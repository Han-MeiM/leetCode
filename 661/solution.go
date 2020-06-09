package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	y_max, x_max := len(M)-1, len(M[0])-1
	if y_max == 0 && x_max == 0 {
		return M
	}

	res := make([][]int, y_max + 1)

	for y, row := range M {
		arr := make([]int, x_max + 1)
		for x := range row {
			x_list, y_list := []int{x}, []int{y}
			if y+1 <= y_max {
				y_list = append(y_list, y+1)
			}
			if y-1 >= 0 {
				y_list = append(y_list, y-1)
			}
			if x+1 <= x_max {
				x_list = append(x_list, x+1)
			}
			if x-1 >= 0 {
				x_list = append(x_list, x-1)
			}
			sum, count := 0, 0
			fmt.Println(x_list, y_list)
			for _, x_one := range x_list {
				for _, y_one := range y_list {
					count++
					sum += M[y_one][x_one]
				}
			}
			arr[x] = sum/count
		}
		res[y] = arr
	}
	return res
}

func main()  {
	res := imageSmoother([][]int{{2,3,4},{5,6,7},{8,9,10},{11,12,13},{14,15,16}})
	fmt.Println(res)
}
