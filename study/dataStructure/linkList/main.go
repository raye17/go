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
	head := new(ListNode)
	head.id = 0
	head.name = "raye"
	node01 := NewListNode(1, "node01")
	head.next = node01
	node02 := NewListNode(2, "node02")
	node01.next = node02
	node03 := NewListNode(3, "node03")
	addNode(head, node03)
	node04 := NewListNode(4, "node04")
	node05 := NewListNode(5, "node05")
	node06 := NewListNode(6, "node06")
	addNodes(head, node04, node05, node06)
	listNode(head)

}
func listNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.id, head.name)
		head = head.next
	}
}
func addNode(head, node *ListNode) {
	for head.next != nil {
		head = head.next
	}
	head.next = node
}
func addNodes(head *ListNode, node ...*ListNode) {
	for _, v := range node {
		addNode(head, v)
	}
}
