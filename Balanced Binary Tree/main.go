package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Abs(integer int) int {
	if integer < 0 {
		return integer * -1
	} else {
		return integer
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isBalanced(root *TreeNode) bool {
	result := recursion(root)
	if result == -1 {
		return false
	} else {
		return true
	}
}

func recursion(head *TreeNode) int {
	if head == nil {
		return 0
	}

	leftDepth := recursion(head.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := recursion(head.Right)
	if rightDepth == -1 {
		return -1
	}

	if Abs(leftDepth-rightDepth) > 1 {
		return -1
	}

	return Max(leftDepth+1, rightDepth+1)
}
