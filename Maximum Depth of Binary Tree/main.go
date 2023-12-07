package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxDepth(root *TreeNode) int {
	n := 0

	if root == nil {
		return n
	}

	return helper(root, n)
}

func helper(root *TreeNode, n int) int {
	n++
	var left int
	var right int

	if root.Left != nil && root.Right != nil {
		left = helper(root.Left, n)
		right = helper(root.Right, n)

		if right > left {
			return right
		} else {
			return left
		}
	}

	if root.Left != nil {
		return helper(root.Left, n)
	}

	if root.Right != nil {
		return helper(root.Right, n)
	}

	return n
}
