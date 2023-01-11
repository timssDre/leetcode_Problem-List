package main

import (
	"fmt"
	"strconv"
)

func main() {
	Request1 := isPalindrome(121)
	Request2 := isPalindrome(-121)
	Request3 := isPalindrome(10)
	fmt.Println(Request1)
	fmt.Println(Request2)
	fmt.Println(Request3)
}

func isPalindrome(x int) bool {
	var request bool = false
	strX := strconv.Itoa(x)
	xLen := len(strX)

	if xLen == 1 {
		return request
	}

	for i := 0; i <= xLen/2; i++ {
		if strX[i] != strX[(xLen-1)-i] {
			return request
		}
	}
	request = true
	return request
}
