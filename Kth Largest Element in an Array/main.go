package main

import "fmt"

func main() {
	test := []int{3, 2, 1, 5, 6, 4}

	fmt.Println(findKthLargest(test, 2))
}

func findKthLargest(nums []int, k int) int {

	return selectKthLargest(nums, 0, len(nums)-1, k)
}

func selectKthLargest(nums []int, left int, right int, k int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	pivotIndex := left

	pivotIndex = partitionSlice(nums, left, right, pivotIndex)

	if len(nums)-k == pivotIndex {
		return nums[pivotIndex]
	} else if len(nums)-k < pivotIndex {
		return selectKthLargest(nums, left, pivotIndex-1, k)
	} else {
		return selectKthLargest(nums, pivotIndex+1, right, k)
	}
}

func partitionSlice(nums []int, left int, right int, pivotIndex int) int {
	pivotValue := nums[pivotIndex]

	// Swapping pivot and the rightmost value
	nums[pivotIndex] = nums[right]
	nums[right] = pivotValue

	storeIndex := left

	for i := left; i < right; i++ {
		if nums[i] <= pivotValue {
			temp := nums[i]
			nums[i] = nums[storeIndex]
			nums[storeIndex] = temp
			storeIndex++
		}
	}

	// Return pivot to its place
	nums[right] = nums[storeIndex]
	nums[storeIndex] = pivotValue

	return storeIndex
}
