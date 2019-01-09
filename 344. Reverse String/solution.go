package main

import "fmt"

//func reverseString(s string) string {
//	var res string
//	for _, v := range s {
//		res = string(v) + res
//	}
//	return res
//}

func reverseString(s string) string {
	l := len(s)
	res := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}

func reverseStringPlanB(s string) string {
	l := len(s)
	res := make([]byte, l)
	for i := 0; i <= l/2; i++ {
		res[i], res[l-i-1] = s[l-i-1], s[i]
	}
	return string(res)
}

func main() {
	str := "asdfgh"
	res := reverseStringPlanB(str)
	fmt.Println(res)
}
