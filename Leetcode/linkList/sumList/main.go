package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
	}
	l1.Next = &ListNode{
		Val: 2,
	}
	for l1 != nil {
		fmt.Println(l1.Val)
	}
}
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var l3 *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			l3 = head
		} else {
			l3.Next = &ListNode{Val: sum}
			l3 = l3.Next
		}
	}
	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}
	return
}
