package main

import "sort"

func main() {
	var nums = []int{-4, -1, 0, 3, 10}
	sortedSquares(nums)
}

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i] = nums[i] * nums[i]
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
