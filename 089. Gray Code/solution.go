package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		res := []int{0}
		return res
	}
	res := grayCode(n - 1)
	resSize := len(res)

	addN := 1
	for n > 1 {
		addN = addN << 1
		n--
	}

	for i := resSize-1; i>=0; i-- {
		res = append(res, addN + res[i])
	}
	return res
}

func main() {
	res := grayCode(4)
	fmt.Println(res)
}
