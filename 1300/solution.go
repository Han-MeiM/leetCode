package main

import "fmt"

func findBestValue(arr []int, target int) int {
	avg := target / len(arr)
	bigArr := make([]int, 0)
	nextTarget := target
	max := 0

	for _, num := range arr {
		if num > avg {
			bigArr = append(bigArr, num)
		} else {
			nextTarget -= num
		}
		if num > max {
			max = num
		}
	}

	if len(arr) == len(bigArr) {
		if abs(avg*len(arr)-target) <= abs((avg+1)*len(arr)-target) {
			return avg
		}
		return avg + 1
	}

	if len(bigArr) == 0 {
		return max
	}

	return findBestValue(bigArr, nextTarget)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func main() {
	arr := []int{4, 9, 3}
	target := 10
	result := findBestValue(arr, target)

	fmt.Println(result)
}
