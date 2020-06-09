package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums3 := []int{}
	i := 0
	j := 0
	for i < m || j < n {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				nums3 = append(nums3, nums1[i])
				i++
			} else {
				nums3 = append(nums3, nums2[j])
				j++
			}
		} else if i < m {
			nums3 = append(nums3, nums1[i:m]...)
			break
		} else if j < n {
			nums3 = append(nums3, nums2[j:]...)
			break
		}
	}
	for k := range nums3 {
		nums1[k] = nums3[k]
	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
