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
	node01 := new(ListNode)
	node01.id = 1
	node01.name = "raye"
	node02 := new(ListNode)
	node02.id = 2
	node02.name = "sxy"
	node01.next = node02
	for {
		if node01 != nil {
			fmt.Println(node01.id, node01.name)
			node01 = node01.next
		}
	}
}
