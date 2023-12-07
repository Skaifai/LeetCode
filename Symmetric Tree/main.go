package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root.Right == nil && root.Left == nil {
		return true
	}

	if root.Left == nil && root.Right != nil {
		return false
	}

	if root.Right == nil && root.Left != nil {
		return false
	}

	leftArray := make([]*TreeNode, 0)
	rightArray := make([]*TreeNode, 0)

	leftArray = leftTreeSearch(root.Left, leftArray)
	rightArray = rightTreeSearch(root.Right, rightArray)

	if len(leftArray) != len(rightArray) {
		return false
	}

	for i := 0; i < len(leftArray); i++ {
		if leftArray[i] != nil && rightArray[i] != nil {
			if leftArray[i].Val != rightArray[i].Val {
				return false
			}
		} else if leftArray[i] != rightArray[i] {
			return false
		}
	}

	return true
}

func leftTreeSearch(root *TreeNode, visited []*TreeNode) []*TreeNode {
	visited = append(visited, root)

	if root.Left != nil {
		visited = leftTreeSearch(root.Left, visited)
	} else {
		visited = append(visited, nil)
	}

	if root.Right != nil {
		visited = leftTreeSearch(root.Right, visited)
	} else {
		visited = append(visited, nil)
	}
	return visited
}

func rightTreeSearch(root *TreeNode, visited []*TreeNode) []*TreeNode {
	visited = append(visited, root)

	if root.Right != nil {
		visited = rightTreeSearch(root.Right, visited)
	} else {
		visited = append(visited, nil)
	}

	if root.Left != nil {
		visited = rightTreeSearch(root.Left, visited)
	} else {
		visited = append(visited, nil)
	}
	return visited
}
