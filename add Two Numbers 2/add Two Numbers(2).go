package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	//l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	//
	////l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	////l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	l1 := ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var i1 string
	var i2 string

	cal := sumPointValue(l1, &i1) + sumPointValue(l2, &i2)
	fmt.Println(cal)
	answer := formingAnswer(cal)
	return answer
}

func formingAnswer(cal int) *ListNode {
	var answer = ListNode{}
	strCal := strconv.Itoa(cal)
	i := utf8.RuneCountInString(strCal)
	i--
	recursion(&answer, &i, strCal)
	return &answer
}

func recursion(answer *ListNode, i *int, strCal string) {
	val, err := strconv.Atoi(string(strCal[*i]))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	answer.Val = val
	*i--
	if *i >= 0 {
		answer.Next = &ListNode{}
		recursion(answer.Next, i, strCal)
	}
}

func sumPointValue(l *ListNode, i *string) int {
	*i = strconv.Itoa(l.Val) + *i
	empty := checkNext(l)
	if empty == true {

		answer, _ := strconv.ParseInt(*i, 10, 64)
		//answer, err := strconv.Atoi(*i)
		//if err != nil {
		//	log.Println(err)
		//}
		//strs := strings.Split(*i, ",")
		//var answer []float64
		//for _, s := range strs {
		//	num, err := strconv.ParseFloat(s) //Atoi(s)
		//	if err == nil {
		//		answer = append(answer, num)
		//	}
		//}
		fmt.Println(answer)
		return 1
	} else {
		return sumPointValue(l.Next, i)
	}
}

func Pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func checkNext(l *ListNode) bool {
	if l.Next == nil {
		//l.Next = &ListNode{}
		return true
	}
	return false
}
