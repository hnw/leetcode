package main

import (
	"fmt"
)

func findMedianSortedArrayAux(nums []int, medianI int, medianNum int) float64 {
	if medianNum == 1 {
		return float64(nums[medianI])
	} else {
		return float64(nums[medianI]+nums[medianI+1]) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	median := 0.0
	i1 := 0
	i2 := 0
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	mergedLen := nums1Len + nums2Len
	medianI := (mergedLen - 1) / 2
	medianNum := 1
	if mergedLen%2 == 0 {
		medianNum = 2
	}

	if nums1Len == 0 {
		return findMedianSortedArrayAux(nums2, medianI, medianNum)
	}
	if len(nums2) == 0 {
		return findMedianSortedArrayAux(nums1, medianI, medianNum)
	}

	// skip to median
	for i := 0; i < medianI; i++ {
		if nums1[i1] <= nums2[i2] {
			i1++
			if i1 >= nums1Len {
				return findMedianSortedArrayAux(nums2, medianI-nums1Len, medianNum)
			}
		} else {
			i2++
			if i2 >= nums2Len {
				return findMedianSortedArrayAux(nums1, medianI-nums2Len, medianNum)
			}
		}
	}

	if nums1[i1] <= nums2[i2] {
		median = float64(nums1[i1])
		i1++
		if i1 >= nums1Len && medianNum == 2 {
			return (median + float64(nums2[i2])) / 2
		}
	} else {
		median = float64(nums2[i2])
		i2++
		if i2 >= nums2Len && medianNum == 2 {
			return (median + float64(nums1[i1])) / 2
		}
	}
	if medianNum == 2 {
		if nums1[i1] <= nums2[i2] {
			median += float64(nums1[i1])
		} else {
			median += float64(nums2[i2])
		}
		median /= 2
	}

	return median
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5, 6}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 11, 14}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5, 6, 7}))
}
