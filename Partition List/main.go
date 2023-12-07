package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//node5 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//node4 := &ListNode{
	//	Val:  5,
	//	Next: node5,
	//}
	//node3 := &ListNode{
	//	Val:  1,
	//	Next: node4,
	//}
	//node2 := &ListNode{
	//	Val:  3,
	//	Next: node3,
	//}

	node1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node0 := &ListNode{
		Val:  2,
		Next: node1,
	}
	partition(node0, 2)
	fmt.Println(node0)
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var previousPointer *ListNode
	currentPointer := head

	for currentPointer != nil && currentPointer.Val < x {
		previousPointer = currentPointer
		currentPointer = currentPointer.Next
	}

	lastSmallerElement := previousPointer

	if currentPointer == nil {
		return head
	}

	for currentPointer != nil {
		if currentPointer.Val < x {
			if lastSmallerElement != nil {
				temp := lastSmallerElement.Next
				lastSmallerElement.Next = currentPointer
				lastSmallerElement = lastSmallerElement.Next
				previousPointer.Next = currentPointer.Next
				currentPointer.Next = temp
				currentPointer = previousPointer
			} else {
				previousPointer.Next = currentPointer.Next
				currentPointer.Next = head
				lastSmallerElement = currentPointer
				head = lastSmallerElement
				currentPointer = previousPointer
			}
		}
		previousPointer = currentPointer
		currentPointer = currentPointer.Next
	}

	return head
}
