package main

func main() {
	//var nums = []int{1, 0, 13, 12, 0}
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
