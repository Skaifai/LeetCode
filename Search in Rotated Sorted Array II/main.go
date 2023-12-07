package main

import "fmt"

func main() {
	test := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(test, 0))
}

func search(nums []int, target int) bool {
	if len(nums) == 1 {
		return target == nums[0]
	}
	result := false
	pivot := pivotSearch(nums)
	if target >= nums[0] {
		result = binarySearch(nums[0:pivot], target)
	} else {
		result = binarySearch(nums[pivot:], target)
	}

	return result
}

func binarySearch(nums []int, target int) (r bool) {
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
			return true
		}
	}
	return false
}

func pivotSearch(nums []int) (r int) {
	if nums[0] == nums[len(nums)-1] {
		i := len(nums) - 1
		for i > 1 && nums[i] >= nums[i-1] {
			i--
		}
		return i
	}

	leftBound := 1
	rightBound := len(nums) - 1
	var pivotIndex int
	for leftBound <= rightBound {
		pivotIndex = (rightBound + leftBound) / 2
		if nums[pivotIndex] >= nums[0] {
			for pivotIndex < len(nums)-1 && nums[pivotIndex] == nums[pivotIndex+1] {
				pivotIndex++
			}
			leftBound = pivotIndex + 1
		} else if nums[pivotIndex] < nums[0] {
			for pivotIndex > 1 && nums[pivotIndex] == nums[pivotIndex-1] {
				pivotIndex--
			}
			rightBound = pivotIndex - 1
		}
	}
	return leftBound
}
