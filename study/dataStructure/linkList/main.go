package main

import "fmt"

type Data struct {
}
type ListNode struct {
	id   int
	name string
	num  int
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
	head.num = 17
	node01 := new(ListNode)
	node02 := new(ListNode)
	node03 := new(ListNode)
	node04 := new(ListNode)
	node05 := new(ListNode)
	node06 := new(ListNode)
	node01.num = 14
	node02.num = 15
	node03.num = 19
	node04.num = 15
	node05.num = 190
	node06.num = 17
	addNodes(head, node01, node02, node03, node04, node05, node06)
	if isPalindrome(head) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
func listNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.id, head.name, head.num)
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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.next == nil {
		return true
	}
	n1, n2 := head, head
	for n2.next != nil && n2.next.next != nil {
		n1 = n1.next      //n1->mid
		n2 = n2.next.next //n2->end
	}
	n2 = n1.next //n2->right part first node
	n1.next = nil
	var n3 *ListNode
	for n2 != nil { //right part convert
		n3 = n2.next
		n2.next = n1
		n1 = n2
		n2 = n3
	}
	n3 = n1
	n2 = head
	for n1 != nil && n2 != nil {
		if n1.num != n2.num {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}
	n1 = n3.next
	n3.next = nil
	for n1 != nil {
		n2 = n1.next
		n1.next = n3
		n3 = n1
		n1 = n2
	}
	return true
}
