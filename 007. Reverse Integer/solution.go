package main

import "fmt"

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func reverse(x int) int {
	var res = 0
	for x != 0 {
		res = res*10 + x % 10
		x = x / 10
	}
	if res > INT_MAX || res < INT_MIN {
		return 0
	} else {
		return res
	}
	return res
}

func main() {
	x := -1534236469
	res := reverse(x)
	fmt.Println(res)
}
