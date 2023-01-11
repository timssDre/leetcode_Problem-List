package main

import "fmt"

func main() {
	var nums []int = []int{5, 7, 7, 8, 8, 10}
	//var nums []int = []int{1}
	//var nums []int = []int{1, 3}
	//var nums []int = []int{2, 2}
	//var nums []int = []int{2, 2, 2}
	answer := searchRange(nums, 8)
	fmt.Println(answer)
}

func searchRange(nums []int, target int) []int {
	var arrAnswer []int
	for i, value := range nums {
		if value == target {
			switch len(arrAnswer) {
			case 0:
				arrAnswer = append(arrAnswer, i)
			case 1:
				arrAnswer = append(arrAnswer, i)
			case 2:
				arrAnswer[1] = i
			}
		}
	}
	if len(arrAnswer) == 1 {
		arrAnswer = append(arrAnswer, arrAnswer[0])
		return arrAnswer
	} else if len(arrAnswer) == 2 {
		return arrAnswer
	}
	arrAnswer = append(arrAnswer, -1)
	arrAnswer = append(arrAnswer, -1)
	return arrAnswer
}
