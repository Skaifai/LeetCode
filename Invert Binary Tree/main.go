package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	right := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root := &TreeNode{
		Val:   2,
		Left:  left,
		Right: right,
	}

	newRoot := invertTree(root)

	fmt.Println(newRoot.Val)
}

func invertTree(root *TreeNode) *TreeNode {
	// Base case
	if root == nil {
		return nil
	}

	//if root.Right == nil && root.Left == nil {
	//	return root
	//}

	////If the binary tree is not full
	//if root.Right == nil {
	//	root.Right = invertTree(root.Left)
	//	root.Left = nil
	//	return root
	//}
	//if root.Left == nil {
	//	root.Left = invertTree(root.Right)
	//	root.Right = nil
	//	return root
	//}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
