package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	recursion(root, &res)
	return res
}

func recursion(head *TreeNode, list *[]int) {
	if head == nil {
		return
	}
	recursion(head.Left, list)
	*list = append(*list, head.Val)
	recursion(head.Right, list)
}
