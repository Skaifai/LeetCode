package main

import "math"

func main() {
	slice := []int{4, 2, 5, 9, 7, 4, 8}
	print(maxProductDifference(slice))
}

func maxProductDifference(nums []int) int {
	firstLargest := math.MinInt
	firstLargestIndex := -1
	secondLargest := math.MinInt
	secondLargestIndex := -1
	firstSmallest := math.MaxInt
	secondSmallest := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] >= firstLargest {
			secondLargest = firstLargest
			firstLargest = nums[i]
			secondLargestIndex = firstLargestIndex
			firstLargestIndex = i
		} else if nums[i] >= secondLargest {
			secondLargest = nums[i]
			secondLargestIndex = i
		}
	}

	nums = append(nums[:firstLargestIndex], nums[firstLargestIndex+1:]...)

	if secondLargestIndex < len(nums) {
		if nums[secondLargestIndex] != secondLargest {
			nums = append(nums[:secondLargestIndex-1], nums[secondLargestIndex-1+1:]...)
		} else {
			nums = append(nums[:secondLargestIndex], nums[secondLargestIndex+1:]...)
		}
	} else {
		nums = nums[:len(nums)-1]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] <= firstSmallest {
			secondSmallest = firstSmallest
			firstSmallest = nums[i]
		} else if nums[i] <= secondSmallest {
			secondSmallest = nums[i]
		}
	}

	return firstLargest*secondLargest - firstSmallest*secondSmallest
}

//} else if nums[i] <= firstSmallest {
//secondSmallest = firstSmallest
//firstSmallest = nums[i]
//} else if nums[i] <= secondSmallest {
//secondSmallest = nums[i]
//}
