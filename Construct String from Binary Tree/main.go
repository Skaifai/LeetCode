package main

import (
	"strconv"
)

//type (
//	Stack struct {
//		top *node
//		length int
//	}
//	node struct {
//		value interface{}
//		prev *node
//	}
//)
//// Create a new stack
//func New() *Stack {
//	return &Stack{nil,0}
//}
//// Return the number of items in the stack
//func (this *Stack) Len() int {
//	return this.length
//}
//// View the top item on the stack
//func (this *Stack) Peek() interface{} {
//	if this.length == 0 {
//		return nil
//	}
//	return this.top.value
//}
//// Pop the top item of the stack and return it
//func (this *Stack) Pop() interface{} {
//	if this.length == 0 {
//		return nil
//	}
//
//	n := this.top
//	this.top = n.prev
//	this.length--
//	return n.value
//}
//// Push a value onto the top of the stack
//func (this *Stack) Push(value interface{}) {
//	n := &node{value,this.top}
//	this.top = n
//	this.length++
//}

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := strconv.Itoa(root.Val)

	if root.Left == nil {
		if root.Right != nil {
			result += "()"
		} else {

		}
	} else {
		result += "(" + tree2str(root.Left) + ")"
	}

	if root.Right == nil {

	} else {
		result += "(" + tree2str(root.Right) + ")"
	}

	return result
}
