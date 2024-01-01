package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	child := 0

	for j := 0; j < len(s) && child < len(g); j++ {
		if s[j] >= g[child] {
			child++
		}
	}

	return child
}

//func findContentChildren(g []int, s []int) int {
//	if len(s) == 0 {
//		return 0
//	}
//
//	sort.Ints(g)
//	sort.Ints(s)
//
//	child := len(g) - 1
//	cookie := len(s) - 1
//
//	for child >= 0 && cookie >= 0 {
//		if s[cookie] >= g[child] {
//			cookie--
//			child--
//			continue
//		}
//		child--
//
//	}
//
//	if cookie == -1 {
//		return len(s)
//	}
//
//	return len(s) - cookie - 1
//}

//func findContentChildren(g []int, s []int) int {
//	sort.Ints(g)
//
//	sort.Ints(s)
//
//	child := len(g) - 1
//
//	for j := len(s) - 1; j >= 0 && child >= 0; j-- {
//		if s[j] >= g[child] {
//			child--
//		}
//	}
//
//	return len(g) - child - 1
//}
