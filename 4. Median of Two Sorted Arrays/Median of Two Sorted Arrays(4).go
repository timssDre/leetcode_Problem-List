package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		//nums1 = []int{1, 3}
		//nums2 = []int{2}

		//nums1 = []int{1, 2}
		//nums2 = []int{3, 4}

		//nums1 = []int{1, 3}
		//nums2 = []int{2, 7}

		//nums1 = []int{1, 3}
		//nums2 = []int{2}

		//nums1 = []int{1, 2}
		//nums2 = []int{4}

		//nums1 = []int{}
		//nums2 = []int{1}

		nums1 []int
		nums2 = []int{2, 3}
	)

	answer := findMedianSortedArrays(nums1, nums2)
	fmt.Println(answer)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lNums1 := len(nums1)
	lNums2 := len(nums2)
	var arr []float64
	for i := 0; i < lNums1; i++ {
		arr = append(arr, float64(nums1[i]))
	}
	for i := 0; i < lNums2; i++ {
		arr = append(arr, float64(nums2[i]))
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	lArr := len(arr)
	if lArr == 1 {
		return arr[0]
	}
	even := lArr%2 == 0
	if even == true {
		middle := lArr / 2
		min := arr[middle-1]
		max := arr[middle]
		answer := (min + max) / 2
		return answer
	} else {
		middle := lArr / 2
		return arr[middle]
	}
}
