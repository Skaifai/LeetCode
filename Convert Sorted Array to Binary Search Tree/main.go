package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	array := []int{-10, -3, 0, 5, 9}

	head := sortedArrayToBST(array)

	print(head)
}

func sortedArrayToBST(nums []int) *TreeNode {
	head := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	recursion(nums, head)

	return head
}

func recursion(nums []int, head *TreeNode) *TreeNode {
	arrayLength := len(nums)

	if arrayLength == 0 {
		return head
	}

	if head == nil {
		head = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}

	middleNum := nums[arrayLength/2]

	head.Val = middleNum

	leftPart := nums[0 : arrayLength/2]

	rightPart := nums[(arrayLength/2)+1 : arrayLength]

	head.Left = recursion(leftPart, head.Left)

	head.Right = recursion(rightPart, head.Right)

	return head
}
