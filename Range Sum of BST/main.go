package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	result := rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high) + root.Val

	return result
}
