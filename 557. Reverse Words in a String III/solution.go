package main

import "fmt"

func reverseWords(s string) string {
	l := len(s)
	word := make([]byte, 0)
	var res string
	for i := l - 1; i >= 0; i-- {
		if s[i] == 32 {
			res = " " + string(word) + res
			word = make([]byte, 0)
			continue
		}
		word = append(word, s[i])
	}
	res = string(word) + res
	return res
}

func main() {
	str := "Let's take LeetCode contest"
	fmt.Println(reverseWords(str))
}
