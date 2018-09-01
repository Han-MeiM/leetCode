package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	res := 0
	for index := range s {
		sum := 1
		for r_index, right := range s[index+1:] {
			if in_str(string(s[index:r_index + index + 1]), right) {
				break
			} else {
				sum += 1
			}
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

func in_str(s string, one int32) bool {
	result := false
	for _, v := range s {
		if one == v {
			result = true
			break
		}
	}
	return result
}

func imooc_programme(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		// 当该字符串最后出现的位置在 start 和 i 之间时
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		// 判断当前子串长度是否大于之前的最长子串长度
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	str := "abcabcbb"
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
	fmt.Println(imooc_programme(str))
}
