package main

func main() {
	t2 := TreeNode{
		Val: 2,
	}
	t3 := TreeNode{
		Val: 3,
	}
	t1 := TreeNode{
		Val:   1,
		Left:  &t2,
		Right: &t3,
	}

	t11 := TreeNode{
		Val:   1,
		Left:  &t3,
		Right: &t2,
	}

	leafSimilar(&t1, &t11)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// Traverse the tree until the children are both empty
	leafs1 := make([]int, 0)
	recursion(root1, &leafs1)

	leafs2 := make([]int, 0)
	recursion(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}

func recursion(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}

	recursion(root.Left, leafs)
	recursion(root.Right, leafs)
}
