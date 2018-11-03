package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	res := true
	var prefix string
	for res {
		for _, v := range strs {
			if i > len(v) - 1 || v[i] != strs[0][i] {
				return prefix
			}
		}
		prefix = prefix + string(strs[0][i])
		i++
	}
	return prefix
}

func main() {
	strs := []string{"aa","a"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
