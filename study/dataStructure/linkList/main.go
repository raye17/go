package main

import (
	"fmt"
	"math"
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

// 回环
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
		n1, n2 = n1.next, n2.next
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

// question:给定两个单链表，判断两个链表是否相交，如果相交，返回第一个相交点
func getIntersectNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

// 1. 找到链表第一个入环节点，如果无环，返回空
func getLoopNode(head *ListNode) *ListNode {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	n1, n2 := head.next, head.next.next
	//如果无环，返回空
	for n1 != n2 {
		if n2.next == nil || n2.next.next == nil {
			return nil
		}
		n2 = n2.next.next
		n1 = n1.next
	}
	//如果有环，n2从头再走一遍，n1,n2下一次相遇在入环节点
	n2 = head
	for n1 != n2 {
		n1 = n1.next
		n2 = n2.next
	}
	return n1
}

// 2. 如果两个链表都无环，返回第一个相交点，不相交返回空
func noLoop(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1, cur2, n := head1, head2, 0
	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}
	//遍历到尾结点，如果不相等，说明不想交
	if cur1 != cur2 {
		return nil
	}
	//此时两个链表相交,n为两个链表长度差，长链表先走n步
	if n >= 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	n = int(math.Abs(float64(n)))
	for n != 0 {
		n--
		cur1 = cur1.next
	}
	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

// 3. 如果只有一个链表有环，必定不想交；当两个链表都有环时
func bothLoop(head1, loop1, head2, loop2 *ListNode) *ListNode {
	var (
		cur1, cur2 *ListNode
	)
	//入环节点相同
	if loop1 == loop2 {
		cur1, cur2 = head1, head2
		n := 0
		for cur1 != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.next
		}
		if n >= 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}
		n = int(math.Abs(float64(n)))
		for n != 0 {
			n--
			cur1 = cur1.next
		}
		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}
