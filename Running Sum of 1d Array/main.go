package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	r := make([]int, len(nums))

	r[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		r[i] = r[i-1] + nums[i]
	}

	return r
}
