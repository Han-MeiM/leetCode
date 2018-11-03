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
	for i:=l-1;i>=0;i-- {
		res = append(res, s[i])
	}
	return string(res)
}

func main() {
	str := "hello"
	res := reverseString(str)
	fmt.Println(res)
}
