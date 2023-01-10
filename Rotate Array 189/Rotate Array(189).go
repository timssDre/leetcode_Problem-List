package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println(nums, k)
}
func rotate(nums []int, k int) {
	first := nums[:k]
	second := nums[k:]
	fmt.Println(first, second)
}
