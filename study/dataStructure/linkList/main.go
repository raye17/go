package main

import (
	"fmt"
)

type Data struct {
}
type ListNode struct {
	id   int
	name string
	num  int
	rand *ListNode
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
	node01 := new(ListNode)
	node02 := new(ListNode)
	node03 := new(ListNode)
	node04 := new(ListNode)
	node05 := new(ListNode)
	node06 := new(ListNode)
	head.num = 0
	node01.num = 1
	node02.num = 2
	node03.num = 3
	node04.num = 4
	node05.num = 5
	node06.num = 6
	head.rand = node03
	node01.rand = node04
	node02.rand = head
	node03.rand = node05
	node04.rand = node01
	node05.rand = node02
	node06.rand = node04
	addNodes(head, node01, node02, node03, node04, node05, node06)
	fmt.Println(&head)
	head01 := copyListWithRand2(head)
	listNode(head01)
	fmt.Println(&head01)
	fmt.Println(head01.rand.rand.num)
	fmt.Println(">>>>>>")
	listNode(head)
}
func copyListWithRand1(head *ListNode) *ListNode {
	m := make(map[*ListNode]*ListNode)
	cur := head
	//key为原node节点，value为新node节点
	for cur != nil {
		node := new(ListNode)
		node.num = cur.num
		m[cur] = node
		cur = cur.next
	}
	cur = head
	for cur != nil {
		m[cur].next = m[cur.next]
		m[cur].rand = m[cur.rand]
		cur = cur.next
	}
	return m[head]
}
func copyListWithRand2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var next *ListNode = nil
	//1->1'->2->2'->3->3'
	for cur != nil {
		next = cur.next
		node := new(ListNode)
		node.num = cur.num
		cur.next = node
		node.next = next
		cur = next
	}
	cur = head
	var node *ListNode = nil
	//新节点node的rand为原节点rand的next
	for cur != nil {
		next = cur.next.next
		node = cur.next
		if cur.rand != nil {
			node.rand = cur.rand.next
		} else {
			node.rand = nil
		}
		cur = next
	}
	res := head.next
	cur = head
	//新旧链表分离
	for cur != nil {
		next = cur.next.next
		node = cur.next
		cur.next = next
		if next != nil {
			node.next = next.next
		} else {
			node.next = nil
		}
		cur = next
	}
	return res
}
func listPartition(head *ListNode, pivot int) *ListNode {
	var (
		sh   *ListNode = nil // small head
		st   *ListNode = nil // small tail
		eh   *ListNode = nil // equal head
		et   *ListNode = nil // equal tail
		bh   *ListNode = nil // big head
		bt   *ListNode = nil // big tail
		next *ListNode = nil
	)
	for head != nil {
		next = head.next
		head.next = nil
		if head.num < pivot {
			if sh == nil {
				sh = head
				st = head
			} else {
				st.next = head
				st = head
			}
		} else if head.num == pivot {
			if eh == nil {
				eh = head
				et = head
			} else {
				et.next = head
				et = head
			}
		} else {
			if bh == nil {
				bh = head
				bt = head
			} else {
				bt.next = head
				bt = head
			}
		}
		head = next
	}
	if st != nil {
		st.next = eh
		if et == nil {
			et = st
		}
	}
	if et != nil {
		et.next = bh
	}
	if sh != nil {
		return sh
	} else if eh != nil {
		return eh
	} else {
		return bh
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
	m := head
	flag := true
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
			flag = false
			break
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
	head = m
	if flag {
		return true
	} else {
		return false
	}
}
