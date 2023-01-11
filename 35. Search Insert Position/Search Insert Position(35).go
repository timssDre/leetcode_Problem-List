package main

import "fmt"

func main() {
	var (
		//nums   = []int{1, 3, 5, 6}
		//target = 5 //2
		//nums   = []int{1, 3, 5, 6}
		//target = 2 //2

		nums   = []int{1, 3, 5, 6}
		target = 7 //2
	)
	result := searchInsert(nums, target)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	for i, value := range nums {
		if nums[i] == target || value > target {
			return i
		}
	}
	return len(nums)
}
