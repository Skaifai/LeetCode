package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP (memoization) + Recursion

func lengthOfLIS(nums []int) int {
	result := 1

	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if dp[i] == 0 {
			dp[i] = recursion(i, nums, dp)
		}
		result = max(result, dp[i])
	}

	return result
}

func recursion(index int, nums []int, dp []int) int {
	if dp[index] != 0 {
		return dp[index]
	}
	res := 0
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			res = max(res, recursion(i, nums, dp))
		}
	}
	dp[index] = res + 1
	return dp[index]
}

// Recursion with no DP

//func lengthOfLIS(nums []int) int {
//	result := 1
//
//	for i := 0; i < len(nums); i++ {
//		result = max(result, recursion(i, nums))
//	}
//
//	return result
//}
//
//func recursion(index int, nums []int) int {
//	res := 0
//	for i := index + 1; i < len(nums); i++ {
//		if nums[i] > nums[index] {
//			res = max(res, recursion(i, nums))
//		}
//	}
//	res++
//	return res
//}
