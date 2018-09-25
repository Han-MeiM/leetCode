package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	d := 1
	for x/d >= 10 {
		d *= 10
	}
	for x > 0 {
		// 最左边的值
		q := x/d
		// 最右边的值
		r := x%10
		if q != r {
			return false
		}
		// x%d 为去掉最左边的值
		// x%d/10 为去掉最右的值
		x = x%d/10
		d /= 100
	}
	return true
}

func main() {
	x := 123321
	res := isPalindrome(x)
	fmt.Println(res)
}
