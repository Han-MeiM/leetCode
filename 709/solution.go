package main

import "fmt"

func toLowerCase(str string) string {
	var res string
	for _, v := range str {
		if 65 <= v && v <= 90 {
			v = v + 32
		}
		res = res + string(v)
	}
	return res
}

func main() {
	str := "Hello"
	res := toLowerCase(str)
	fmt.Println(res)
}
