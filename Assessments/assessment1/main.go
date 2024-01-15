package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))

	addTwoNumbers(&ListNode{
		Val:  5,
		Next: nil,
	}, &ListNode{
		Val:  5,
		Next: nil,
	})

	//addTwoNumbers(&ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}, &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//})

	//addTwoNumbers(&ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val: 9,
	//				Next: &ListNode{
	//					Val: 9,
	//					Next: &ListNode{
	//						Val:  9,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}, &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val:  9,
	//				Next: nil,
	//			},
	//		},
	//	},
	//})
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 >= len(s) {
				res++
				break
			}
			if s[i+1] == 'V' {
				i++
				res += 4
			} else if s[i+1] == 'X' {
				i++
				res += 9
			} else {
				res++
			}
		case 'V':
			res += 5
		case 'X':
			if i+1 >= len(s) {
				res += 10
				break
			}
			if s[i+1] == 'L' {
				i++
				res += 40
			} else if s[i+1] == 'C' {
				i++
				res += 90
			} else {
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			if i+1 >= len(s) {
				res += 100
				break
			}
			if s[i+1] == 'D' {
				i++
				res += 400
			} else if s[i+1] == 'M' {
				i++
				res += 900
			} else {
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}

	}
	return res
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

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	num1 := 0
//	num2 := 0
//	multiplier := 1
//	for l1 != nil {
//		num1 += multiplier * l1.Val
//		multiplier *= 10
//		l1 = l1.Next
//	}
//	multiplier = 1
//	for l2 != nil {
//		num2 += multiplier * l2.Val
//		multiplier *= 10
//		l2 = l2.Next
//	}
//	println(num1)
//	println(num2)
//	sum := num1 + num2
//	println(sum)
//	res := &ListNode{
//		Val:  sum % 10,
//		Next: nil,
//	}
//	track := res
//	sum /= 10
//	for sum != 0 {
//		chain := &ListNode{
//			Val:  sum % 10,
//			Next: nil,
//		}
//		sum /= 10
//		track.Next = chain
//		track = chain
//	}
//	return res
//}
