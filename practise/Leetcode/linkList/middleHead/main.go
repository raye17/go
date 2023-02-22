package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func middleNode(head *ListNode) *ListNode {
	var quick, slow = head, head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}
