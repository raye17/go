package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head01 := &ListNode{1, nil}
	head02 := &ListNode{2, nil}
	head03 := &ListNode{3, nil}
	head04 := &ListNode{4, nil}
	head05 := &ListNode{5, nil}
	head01.Next = head02
	head02.Next = head03
	head03.Next = head04
	head04.Next = head05
	//removeNthFromEnd(head01,2)
	var head = head01
	ret := removeNthFromEnd(head, 2)
	for ret != nil {
		fmt.Println(ret)
		ret = ret.Next
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		count++
	}
	dummy := &ListNode{0, head}
	cur := dummy
	edit := count - n
	for i := 0; i < edit; i++ {
		fmt.Println("cur:", cur)
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
func removeNthFromEnd01(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	second, right := dummy, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for ; right != nil; right = right.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
