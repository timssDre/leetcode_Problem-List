// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {

	nums := []int{2, 1, 9, 4, 4, 56, 90, 3}
	answer := twoSum(nums, 8)
	fmt.Println(answer)

	//nums := []int{-18, 12, 3, 0}
	//answer := twoSum(nums, -6)
	//fmt.Println(answer)

	//nums := []int{-3, 4, 3, 90}
	//answer := twoSum(nums, -0)
	//fmt.Println(answer)

	//nums := []int{-10, -1, -18, -19}
	//answer := twoSum(nums, -19)
	//fmt.Println(answer)

	//nums := []int{-1, -2, -3, -4, -5}
	//answer := twoSum(nums, -8)
	//fmt.Println(answer)

	//nums := []int{2, 7, 11, 15}
	//answer := twoSum(nums, 9)
	//fmt.Println(answer)
	//
	//nums1 := []int{3, 2, 4}
	//answer1 := twoSum(nums1, 6)
	//fmt.Println(answer1)
	//
	//nums2 := []int{3, 3}
	//answer2 := twoSum(nums2, 6)
	//fmt.Println(answer2)
}

//func twoSum(nums []int, target int) []int {
//	arrAnswer := make([]int, 0)
//	calculations := 0
//	for i, value := range nums {
//		if target <= 0 {
//			if value == 0 {
//				arrAnswer = append(arrAnswer, i)
//			} else if value < 0 {
//				if calculations != 0 && calculations == target {
//					break
//				}
//				arrAnswer = append(arrAnswer, i)
//				calculations = value
//				arrAnswer = make([]int, 0)
//				arrAnswer = append(arrAnswer, i)
//				for j, k := range nums {
//					if j == i || calculations+k != target {
//						continue
//					}
//					arrAnswer = append(arrAnswer, j)
//					calculations += k
//					break
//				}
//				if calculations == target {
//					break
//				}
//			}
//			continue
//		} else if calculations >= target {
//			break
//		} else if value > target {
//			continue
//		}
//		calculations = 0
//		calculations += value
//		arrAnswer = make([]int, 0)
//		arrAnswer = append(arrAnswer, i)
//
//		for j, k := range nums {
//			if j == i || calculations+k != target {
//				continue
//			}
//			arrAnswer = append(arrAnswer, j)
//			calculations += k
//			break
//		}
//	}
//
//	return arrAnswer
//}

func twoSum(nums []int, target int) []int {
	var output = []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i)
				output = append(output, j)
			}
		}
	}
	return output
}
