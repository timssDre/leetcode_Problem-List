package main

func main() {
	//var nums = []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3

	//var nums = []int{-1}
	//k := 2

	//var nums = []int{1, 2}
	//k := 3

	var nums = []int{1}
	k := 1

	//var nums = []int{1, 2}
	//k := 2

	//var nums = []int{1, 2, 3}
	//k := 4

	//var nums = []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3

	rotate(nums, k)
}

func rotate(nums []int, k int) {
	lNums := len(nums)
	if lNums == 1 {
		return
	} else if lNums <= k {
		aDifference := lNums - k
		a := k - (1 - aDifference)
		recursion(nums, a, lNums)
		k -= a
		rotate(nums, k)
	} else {
		recursion(nums, k, lNums)
	}
}
func recursion(nums []int, k int, lNums int) {
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

//альтернатива
//import "fmt"
//func rotate(nums []int, k int)  {
//	m := k % len(nums)
//	copy(nums, append(nums[len(nums)-m:],nums[:len(nums)-m]...))
//}
