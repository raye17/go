package main

import (
	"fmt"
	"math"
)

// BinaryTree 二叉树结构
type BinaryTree struct {
	value       interface{}
	left, right *BinaryTree
}

func (bt *BinaryTree) setValue(v interface{}) {
	bt.value = v
}

// Stack 栈
type Stack struct {
	array []*BinaryTree
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.array = []*BinaryTree{}
	return stack
}
func (s *Stack) Push(bt *BinaryTree) {
	s.array = append(s.array, bt)
}
func (s *Stack) Pop() *BinaryTree {
	item := s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]
	return item
}
func (s *Stack) size() int {
	return len(s.array)
}
func (s *Stack) IsEmpty() bool {
	return len(s.array) == 0
}

type Set []interface{}

// CreateTree 数组建立二叉树递归实现
func CreateTree(arr []interface{}, i int) *BinaryTree {
	node := &BinaryTree{arr[i], nil, nil}
	if 2*i < len(arr) && arr[2*i] != nil {
		node.left = CreateTree(arr, 2*i)
	}
	if 2*i+1 < len(arr) && arr[2*i+1] != nil {
		node.right = CreateTree(arr, 2*i+1)
	}
	return node
}
func NewBT(v interface{}) *BinaryTree {
	bt := &BinaryTree{}
	bt.setValue(v)
	return bt
}

var set Set

func main() {
	//head := &BinaryTree{}
	//head.setValue("001")
	//node01 := NewBT("002")
	//node02 := NewBT("node02")
	//node03 := NewBT(3)
	//node04 := NewBT(4)
	//node05 := NewBT("a")
	//node06 := NewBT("b")
	//node07 := NewBT("c")
	//node08 := NewBT("d")
	//head.left = node01
	//head.right = node02
	//node01.left = node03
	//node01.right = node04
	//node02.left = node05
	//node02.right = node06
	//node03.left = node07
	//node03.right = node08
	//arr := []interface{}{nil, "001", "002", "node02", 3, 4, "a", "b", "c", "d"}
	arr1 := []interface{}{nil, 5, 3, 7, 2, 4, 6, 8, 1}
	head := CreateTree(arr1, 1)
	//f(head)
	//fmt.Println(set)
	preOrderUnRecur(head)
	posOrderUnRecur(head)
	inOrderUnRecur(head)
	fmt.Println(IsBST(head))
}

// 递归遍历
func f(head *BinaryTree) {
	if head == nil {
		return
	}
	//fmt.Println("先序遍历: ", head.value)
	set = append(set, head.value, "->")
	f(head.left)
	//fmt.Println("中序:", head.value)
	//set = append(set, head.value,"->")
	f(head.right)
	//fmt.Println("后序遍历:", head.value)
	//set = append(set, head.value, "->")
}

// 非递归遍历
/*
先序遍历：
1.先把head节点压入栈中
2.从栈中弹出一个节点
3.处理该节点
4.如果有子树，先右后左压入栈中
5.重复 2,3,4 过程
中序遍历：
1.把整棵左子树压入栈
2.把节点弹出栈，判断该节点的右子树，重复1
后序遍历：
1.把head节点压入栈中
2.从栈中弹出一个节点，标记为cur节点
3.cur放入另一个栈
4.如果有子树，先左后右
5.重复2,3,4过程
*/
func preOrderUnRecur(head *BinaryTree) {
	set := Set{}
	fmt.Println("pre-order: ")
	if head != nil {
		stack := NewStack()
		stack.Push(head)
		for !stack.IsEmpty() {
			head = stack.Pop()
			set = append(set, head.value, "->")
			if head.right != nil {
				stack.Push(head.right)
			}
			if head.left != nil {
				stack.Push(head.left)
			}
		}
	}
	length := len(set)
	if set[length-1] == "->" {
		set = set[:length-1]
	}
	fmt.Println(set)
}
func inOrderUnRecur(head *BinaryTree) {
	set := Set{}
	fmt.Println("非递归，中序遍历：")
	if head == nil {
		return
	}
	stack01 := NewStack()
	for head != nil || !stack01.IsEmpty() {
		if head != nil {
			stack01.Push(head)
			head = head.left
		} else {
			head = stack01.Pop()
			set = append(set, head.value, "->")
			head = head.right
		}
	}
	length := len(set)
	if set[length-1] == "->" {
		set = set[:length-1]
	}
	fmt.Println(set)
}
func posOrderUnRecur(head *BinaryTree) {
	set := Set{}
	fmt.Println("非递归，后序遍历：")
	if head != nil {
		stack01 := NewStack()
		stack02 := NewStack()
		stack01.Push(head)
		for !stack01.IsEmpty() {
			cur := stack01.Pop()
			stack02.Push(cur)
			if cur.left != nil {
				stack01.Push(cur.left)
			}
			if cur.right != nil {
				stack01.Push(cur.right)
			}
		}
		for !stack02.IsEmpty() {
			cur := stack02.Pop()
			set = append(set, cur.value, "->")
		}
	}
	length := len(set)
	if set[length-1] == "->" {
		set = set[:length-1]
	}
	fmt.Println(set)
}

//二叉树的宽度遍历，队列

// 搜索二叉树：左子树<根节点<右子树;中序遍历的结果应为非递减
// 判断是否为搜索二叉树
var preValue = math.MinInt

func IsBST(node *BinaryTree) bool {
	if node == nil {
		return true
	}
	if !IsBST(node.left) {
		return false
	}
	v := node.value.(int)
	if v <= preValue {
		return false
	} else {
		preValue = v
	}
	return IsBST(node.right)
}
