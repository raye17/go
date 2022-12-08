package main

import "fmt"

type ListNode struct {
	id   int
	name string
	next *ListNode
}

func NewListNode(id int, name string) *ListNode {
	node := new(ListNode)
	node.id = id
	node.name = name
	return node
}
func main() {
	var head *ListNode
	head = &ListNode{0, "raye", nil}
	tail := head
	fmt.Println("head:", head, "*head:", *head, "&head:", &head)
	fmt.Println(&tail)
	var node = &ListNode{2, "test", nil}
	tail.next = node
	fmt.Println(head.next)
	tail = node
	if head.next == tail {
		fmt.Println("1")
	}

	//head = new(ListNode)
	//head.id = 0
	//head.name = "raye"
	//tail := head
	//for i := 1; i <= 3; i++ {
	//	switch i {
	//	case 1:
	//		tail.next = &ListNode{
	//			id:   1,
	//			name: "test01",
	//		}
	//		tail = tail.next
	//	case 2:
	//		tail.next = &ListNode{
	//			id:   2,
	//			name: "test02",
	//		}
	//	case 3:
	//		tail.next = &ListNode{
	//			id:   3,
	//			name: "test03",
	//		}
	//		tail = tail.next
	//	}
	//}
	//fmt.Print(*head)
	//for head.next != nil {
	//
	//	head = head.next
	//	fmt.Print("->", *head)
	//}
}
