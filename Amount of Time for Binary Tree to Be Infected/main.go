package main

import "fmt"

func main() {
	n1 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	//n1 := TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val: 4,
	//				Left: &TreeNode{
	//					Val:   5,
	//					Left:  nil,
	//					Right: nil,
	//				},
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: nil,
	//}
	fmt.Println(amountOfTime(&n1, 3))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaximumDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}

	leftDepth := getMaximumDepth(root.Left) + 1
	rightDepth := getMaximumDepth(root.Right) + 1

	return max(leftDepth, rightDepth)
}

func findInfection(root *TreeNode, start int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == start {
		return root
	}

	left := findInfection(root.Left, start)
	right := findInfection(root.Right, start)

	if left != nil {
		return left
	}

	return right
}

func amountOfTime(root *TreeNode, start int) int {
	maxDepth := getMaximumDepth(root)
	if root.Val == start {
		return maxDepth
	}

	connections := make(map[*TreeNode]*TreeNode)

	rememberParents(root, connections)

	//fmt.Println(connections)

	return count(findInfection(root, start), nil, connections, 0, 0)
}

func rememberParents(root *TreeNode, connections map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		//connections[root] = append(connections[root], root.Left)
		connections[root.Left] = root
	}
	if root.Right != nil {
		//connections[root] = append(connections[root], root.Right)
		connections[root.Right] = root
	}

	rememberParents(root.Left, connections)
	rememberParents(root.Right, connections)
}

func count(current *TreeNode, previous *TreeNode, connections map[*TreeNode]*TreeNode, depth int, before int) int {
	if previous == nil {
		return count(connections[current], current, connections, getMaximumDepth(current), before+1)
	}

	if connections[current] == nil {
		if current.Left != nil && previous == current.Left {
			rightDepth := getMaximumDepth(current.Right)
			if rightDepth+before > depth {
				depth = rightDepth + before
			}
			return depth
		}
		leftDepth := getMaximumDepth(current.Left)
		if leftDepth+before > depth {
			depth = leftDepth + before
		}
		return depth
	}

	if current.Left != nil && previous == current.Left {
		rightDepth := getMaximumDepth(current.Right)
		if rightDepth+before > depth {
			depth = rightDepth + before
		}
		return count(connections[current], current, connections, depth, before+1)
	}

	leftDepth := getMaximumDepth(current.Left)
	if leftDepth+before > depth {
		depth = leftDepth + before
	}
	return count(connections[current], current, connections, depth, before+1)
}

//func rememberParents(root *TreeNode, connections map[*TreeNode]*TreeNode, depths map[*TreeNode]int) int {
//	if root == nil {
//		return 0
//	}
//
//	if root.Left != nil {
//		//connections[root] = append(connections[root], root.Left)
//		connections[root.Left] = root
//	}
//	if root.Right != nil {
//		//connections[root] = append(connections[root], root.Right)
//		connections[root.Right] = root
//	}
//
//	leftDepth := rememberParents(root.Left, connections, depths)
//	rightDepth := rememberParents(root.Right, connections, depths)
//	depths[root] = max(leftDepth, rightDepth)
//	return depths[root] + 1
//}

//func recursion(root *TreeNode, start int) (int, bool) {
//	if root == nil {
//		return 0, false
//	}
//
//	if root.Val == start {
//		return getMaximumDepth(root), true
//	}
//
//	leftDepth, isLeft := recursion(root.Left, start)
//
//	if isLeft {
//		rightDepth := getMaximumDepth(root.Right)
//		return max(leftDepth, rightDepth), true
//	}
//
//	rightDepth, isRight := recursion(root.Right, start)
//	if isRight {
//		return max(leftDepth,rightDepth), true
//	}
//
//	return max(leftDepth, rightDepth), false
//}

//func amountOfTime(root *TreeNode, start int) int {
//	if root == nil {
//		return 0
//	}
//
//	if root.Val == start {
//
//		return 0
//	}
//
//	return 0
//}

//func amountOfTime(root *TreeNode, start int) int {
//	maxDepth := getMaximumDepth(root)
//	if root.Val == start {
//		return maxDepth
//	}
//
//	level, isLeft := findInfection(root.Left, start)
//
//	if isLeft {
//		fmt.Println("Left", level+1)
//	}
//
//	level, isRight := findInfection(root.Right, start)
//
//	if isRight {
//		fmt.Println("Right", level+1)
//	}
//
//	return 0
//}

//func findInfection(root *TreeNode, start int) (int, bool) {
//	if root == nil {
//		return 0, false
//	}
//
//	if root.Val == start {
//		return 0, true
//	}
//
//	leftDepth, isLeft := findInfection(root.Left, start)
//
//	if isLeft {
//		return leftDepth + 1, isLeft
//	}
//
//	rightDepth, isRight := findInfection(root.Right, start)
//
//	if isRight {
//		return rightDepth + 1, isRight
//	}
//
//	return 0, false
//}
