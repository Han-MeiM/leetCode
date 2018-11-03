package two_sum

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, val := range nums {
		diff := target - val
		// j, ok  =  value, bool
		if j, ok := dict[diff]; ok {
			return []int{i, j}
		}
		dict[val] = i
	}
	return nil
}

func main() {
	num := [...]int{11, 7, 2, 15}
	target := 9
	res := twoSum(num[:], target)
	fmt.Println(res)
}
