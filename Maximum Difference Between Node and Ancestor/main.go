package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	fmt.Println(maxAncestorDiff(n1))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	childrenValues := make([]int, 0)
	getChildren(root, &childrenValues)
	res := 0
	for i := 0; i < len(childrenValues); i++ {
		if abs(root.Val-childrenValues[i]) > res {
			res = abs(root.Val - childrenValues[i])
		}
	}
	left := maxAncestorDiff(root.Left)
	right := maxAncestorDiff(root.Right)

	return max(res, max(left, right))
}

func getChildren(root *TreeNode, childrenValues *[]int) {
	if root == nil {
		return
	}

	*childrenValues = append(*childrenValues, root.Val)
	getChildren(root.Left, childrenValues)
	getChildren(root.Right, childrenValues)
}
