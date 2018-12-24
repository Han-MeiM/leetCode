package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := ExpandAroundCenter(s, i, i)
		len2 := ExpandAroundCenter(s, i, i+1)
		length := Max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func ExpandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func main() {
	str := "babad"
	res := longestPalindrome(str)
	fmt.Println(res)
}
