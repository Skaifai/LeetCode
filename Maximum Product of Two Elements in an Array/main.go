package main

import "math"

func main() {
	array := []int{3, 4, 5, 2}
	print(maxProduct(array))
}

func maxProduct(nums []int) int {
	first := math.MinInt
	second := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] >= first {
			second = first
			first = nums[i]
		} else if nums[i] > second {
			second = nums[i]
		}
	}
	return (first - 1) * (second - 1)
}
