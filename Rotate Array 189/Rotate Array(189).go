package main

func main() {
	//var nums = []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3

	//var nums = []int{-1}
	//k := 2

	//var nums = []int{1, 2}
	//k := 2

	var nums = []int{1, 2, 3}
	k := 4
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	lNums := len(nums)
	first := nums[:lNums-k]
	second := nums[lNums-k:]

	seven := make([]int, lNums)
	s := seven[:0]
	s = append(s, second[:]...)
	s = append(s, first[:]...)
	for i, value := range seven {
		nums[i] = value
	}
}
