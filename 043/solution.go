package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	arr := make([]int, n1+n2)
	intNum1, intNum2 := make([]int, n1), make([]int, n2)

	for j := n2 - 1; j >= 0; j-- {
		intNum2[j], _ = strconv.Atoi(string(num2[j]))
	}

	for i := n1 - 1; i >= 0; i-- {
		intNum1[i], _ = strconv.Atoi(string(num1[i]))
		for j := n2 - 1; j >= 0; j-- {
			// 从个位数开始获取单个乘法值
			temp := intNum1[i] * intNum2[j]
			temp = arr[i+j+1] + temp
			arr[i+j+1] = temp % 10
			// 若当前位数的进一位值超过10先放置不管，等下一位乘法来做处理
			arr[i+j] = arr[i+j] + temp/10
		}
	}
	return intArrToString(arr)
}

// 数字数组转字符串
func intArrToString(arr []int) string {
	var res string
	zeroFlag := true
	for _, item := range arr {
		if item == 0 && zeroFlag {
			continue
		}
		if zeroFlag {
			zeroFlag = false
		}
		res = res + strconv.Itoa(item)
	}
	return res
}

func main() {
	num1, num2 := "123", "456"
	res := multiply(num1, num2)
	fmt.Println(res)
}
