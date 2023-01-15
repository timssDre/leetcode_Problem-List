package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//ln := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ln := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	//gfd:= ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: nil,
	//	},
	//}
	middleNode(&ln)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var (
		j   []int
		min int
		max int
	)
	a := size(head, j)
<<<<<<< HEAD

	answer := ListNode{}
	min = len(a) / 2
	max = len(a) - 1

=======
	b := check(len(a))

	answer := ListNode{}
	if b == false {
		min = len(a) / 2
		max = len(a) - 1
	} else {
		min = len(a) / 2
		max = len(a) - 1
	}
>>>>>>> origin/main
	reqAnswer(a, &answer, min, max)

	return &answer
}

func reqAnswer(a []int, answer *ListNode, min int, max int) {
	answer.Val = a[min]
	if min != max {
		answer.Next = &ListNode{}
		min++
		reqAnswer(a, answer.Next, min, max)
	}
}

<<<<<<< HEAD
=======
func check(n int) bool {
	if n < 2 {
		return n%2 == 0
	}
	return check(n - 2)
}

>>>>>>> origin/main
func size(head *ListNode, j []int) []int {
	j = append(j, head.Val)
	if nil != head.Next {
		a := size(head.Next, j)
		return a
	}
	return j
}
