package main

func main() {
	node1 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node1,
	}
	node3 := &ListNode{
		Val:  1,
		Next: node2,
	}

	node4 := &ListNode{
		Val:  5,
		Next: nil,
	}
	//node5 := &ListNode{
	//	Val:  3,
	//	Next: node4,
	//}
	//node6 := &ListNode{
	//	Val:  1,
	//	Next: node5,
	//}

	result := mergeTwoLists(node3, node4)
	println(result)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		originalHead := list2
		for list2.Next != nil && list2.Next.Val < list1.Val {
			list2 = list2.Next
		}
		temp := list2.Next
		list2.Next = list1
		list1.Next = mergeTwoLists(temp, list1.Next)
		return originalHead
	} else {
		originalHead := list1
		for list1.Next != nil && list1.Next.Val < list2.Val {
			list1 = list1.Next
		}
		temp := list1.Next
		list1.Next = list2
		list2.Next = mergeTwoLists(temp, list2.Next)
		return originalHead
	}
}
