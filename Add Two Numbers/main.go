package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	pureSum := 0
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	track := res

	for l1 != nil && l2 != nil {
		pureSum = l1.Val + l2.Val + carry
		carry = pureSum / 10
		if carry != 0 {
			pureSum %= 10
		}
		track.Val = pureSum
		if l1.Next != nil || l2.Next != nil {
			track.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil && carry != 0 {
			track.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
			return res
		}
		track = track.Next
	}

	if l1 == nil {
		l1 = l2
	}

	for l1 != nil {
		pureSum = l1.Val + carry
		carry = pureSum / 10
		if carry != 0 {
			pureSum %= 10
		}
		track.Val = pureSum
		if l1.Next != nil {
			track.Next = &ListNode{}
		} else {
			if carry != 0 {
				track.Next = &ListNode{
					Val:  carry,
					Next: nil,
				}
			}
			break
		}
		track = track.Next
		l1 = l1.Next
	}

	return res
}
