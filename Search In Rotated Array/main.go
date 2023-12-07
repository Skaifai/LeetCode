package main

import "fmt"

func main() {
	test := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(test, 0))
}

func search(nums []int, target int) int {
	result := -1

	if nums[0] > nums[len(nums)-1] {
		pivot := pivotSearch(nums)
		if target >= nums[0] {
			result = binarySearch(nums[0:pivot], target)
		} else {
			intermediateResult := binarySearch(nums[pivot:], target)
			if intermediateResult < 0 {
				result = intermediateResult
			} else {
				result = intermediateResult + pivot
			}
		}
	} else {
		result = binarySearch(nums, target)
	}

	return result
}

func binarySearch(nums []int, target int) (r int) {
	leftBound := 0
	rightBound := len(nums) - 1
	var searchIndex int
	for leftBound <= rightBound {
		searchIndex = (rightBound + leftBound) / 2
		if nums[searchIndex] < target {
			leftBound = searchIndex + 1
		} else if nums[searchIndex] > target {
			rightBound = searchIndex - 1
		} else {
			return searchIndex
		}
	}
	return -1
}

func pivotSearch(nums []int) (r int) {
	leftBound := 1
	rightBound := len(nums) - 1
	var pivotIndex int
	for leftBound <= rightBound {
		pivotIndex = (rightBound + leftBound) / 2
		if nums[pivotIndex] > nums[0] {
			leftBound = pivotIndex + 1
		} else {
			rightBound = pivotIndex - 1
		}
	}
	return leftBound
}
