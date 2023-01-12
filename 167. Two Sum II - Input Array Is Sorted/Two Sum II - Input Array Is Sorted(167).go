package main

import "fmt"

func main() {

	var (
		nums1   = []int{2, 7, 11, 15}
		nums2   = []int{2, 3, 4}
		nums3   = []int{-1, 0}
		target1 = 9
		target2 = 6
		target3 = -1
	)

	nums11 := twoSum(nums1, target1) //[1,2]
	nums22 := twoSum(nums2, target2) //[1,3]
	nums33 := twoSum(nums3, target3) //[1,2]
	fmt.Println(nums11, nums22, nums33)
}

func twoSum(numbers []int, target int) []int {
	var answer []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] <= target {
			for j := i + 1; j < len(numbers); j++ {
				if numbers[i]+numbers[j] == target {
					answer = append(answer, i+1, j+1)
					return answer
				}
			}
		}
	}
	return answer
}

//func twoSum(numbers []int, target int) []int {
//	first := 0;
//	second := len(numbers)-1
//	for first < second {
//		if numbers[first]+numbers[second] == target {
//			break
//		} else if numbers[first]+numbers[second] < target {
//			first +=1
//		} else {
//			second -= 1
//		}
//	}
//	return []int{first+1, second+1}
//}
