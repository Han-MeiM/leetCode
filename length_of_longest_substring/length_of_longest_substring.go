package length_of_longest_substring

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

func main() {
	str := "abcabcbb"
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}
