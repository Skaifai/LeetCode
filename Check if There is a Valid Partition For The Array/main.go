package main

import "fmt"

func main() {
	test := []int{0, 0, 0, 0}
	fmt.Println(validPartition(test))
}

func validPartition(nums []int) bool {

	dp := []bool{true, false, nums[0] == nums[1]}

	recursion(nums, dp, 2)

	return dp[2]
}

func recursion(nums []int, dp []bool, index int) bool {
	if index > len(nums)-1 {
		return false
	}
	currentResult := false
	if nums[index] == nums[index-1] && dp[1] {
		currentResult = true
	} else if nums[index] == nums[index-1] && nums[index-1] == nums[index-2] && dp[0] {
		currentResult = true
	} else if nums[index] == nums[index-1]+1 && nums[index-1] == nums[index-2]+1 && dp[0] {
		currentResult = true
	}

	dp[0] = dp[1]
	dp[1] = dp[2]
	dp[2] = currentResult
	return recursion(nums, dp, index+1)
}
