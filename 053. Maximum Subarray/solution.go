package main

import "fmt"

// O(n)解法
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
}

// 分治
func maxSubSumRec(nums []int, l int, r int) int {
	if l == r { return nums[l] }
	if l == r-1 {
		return max(nums[l], max(nums[r], nums[l]+nums[r]))
	}
	mid := (l+r)>>1
	lret := maxSubSumRec(nums, l, mid-1)
	rret := maxSubSumRec(nums, mid+1, r)

 	mret := nums[mid]
 	sum := mret
	for i := mid-1; i >= l ; i-- {
		sum += nums[i]
		mret = max(mret, sum)
	}

	sum = mret

	for i := mid+1; i <= r; i ++ {
		sum += nums[i]
		mret = max(mret, sum)
	}

	return max(lret, max(rret, mret));
}

func max (num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(arr)
	res2 := maxSubSumRec(arr, 0, len(arr) - 1)
	fmt.Println(res)
	fmt.Println(res2)
}
