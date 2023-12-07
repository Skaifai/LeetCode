package main

import (
	"fmt"
	"math"
)

func enqueue(queue []int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []int) ([]int, int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		queue = []int{}
		return queue, element
	}
	queue = queue[1:]
	return queue, element // Slice off the element once it is dequeued.
}

func peek(queue []int) int {
	return queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	currentMaxIndex := 0
	currentMaxValue := nums[currentMaxIndex]

	for i := 0; i < k; i++ {
		if nums[i] > currentMaxValue {
			currentMaxValue = nums[i]
			currentMaxIndex = i
		}
	}

	result[0] = currentMaxValue

	for i := 1; i < len(result); i++ {
		if i > currentMaxIndex {
			currentMaxValue = math.MinInt
			for j := i; j < i+k; j++ {
				if nums[j] > currentMaxValue {
					currentMaxIndex = j
					currentMaxValue = nums[currentMaxIndex]
				}
			}
			result[i] = currentMaxValue
			continue
		}

		if nums[i+k-1] > currentMaxValue {
			currentMaxIndex = i + k - 1
			currentMaxValue = nums[currentMaxIndex]
			result[i] = currentMaxValue
			continue
		}

		result[i] = currentMaxValue
	}
	return result
}

func main() {
	test := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -10, -20, 120}
	fmt.Println(maxSlidingWindow(test, 3))
}
