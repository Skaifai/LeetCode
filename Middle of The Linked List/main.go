package main

import "fmt"

func main() {
	//link4 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}

	link3 := &ListNode{
		Val:  3,
		Next: nil,
	}

	link2 := &ListNode{
		Val:  2,
		Next: link3,
	}

	link1 := &ListNode{
		Val:  1,
		Next: link2,
	}

	result := middleNode(link1)
	fmt.Println(result)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	anotherLink := head
	for anotherLink != nil && anotherLink.Next != nil {
		head = head.Next
		anotherLink = anotherLink.Next.Next
	}
	return head
}
