package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row := len(nums)
	column := len(nums[0])
	if row * column != r * c {
		return nums
	}

	new_nums := make([]int, row * column)
	

}

func main() {
	
}
