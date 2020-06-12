package main

import "fmt"

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))

	for index, left := range T{
		for offset, right := range T[index:]{
			if left < right {
				result[index] = offset
				break
			}
		}
	}

	return result
}

func main() {
	test := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(test)

	fmt.Println(result)
}
