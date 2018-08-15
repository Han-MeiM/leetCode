package find_median_sorted_arrays

import "fmt"

// 合并两个数组并排序
func mergeOrder(nums1 []int, nums2 []int) []int  {
	length := len(nums1) + len(nums2)
	// 声明最终合并后的数组结果
	result := make([]int, length)
	m, n := 0, 0
	for i := range result {
		if m == len(nums1) {
			result[i] = nums2[n]
			n++
		} else if n == len(nums2) {
			result[i] = nums1[m]
			m++
		} else {
			if nums1[m] < nums2[n] {
				result[i] = nums1[m]
				m++
			} else {
				result[i] = nums2[n]
				n++
			}
		}
	}
	return result[:]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := mergeOrder(nums1, nums2)
	length := len(nums3)
	mid := length / 2
	result := 0.0
	if length % 2 != 0 {
		result = float64(nums3[mid])
	} else {
		result = (float64(nums3[mid]) + float64(nums3[mid-1])) / 2
	}
	return result
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
