package main

import (
	"fmt"
)

//Example 1:
//
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
//Example 2:
//
//Input: nums = [-1,0,3,5,9,12], target = 2
//Output: -1
//Explanation: 2 does not exist in nums so return -1
func main() {
	var nums = []int{-1, 0, 3, 5, 9, 12}
	//var nums = []int{5}
	output := -5
	//var nums = []int{-1, 0, 5}
	//output := 0
	//var nums = []int{-1, 0, 3, 5, 9, 12}
	//output := 2

	answer := search(nums, output)
	fmt.Println(answer)
}

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
