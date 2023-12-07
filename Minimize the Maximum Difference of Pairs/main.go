package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	test := []int{4, 0, 2, 1, 2, 5, 5, 3}
	fmt.Println(minimizeMax(test, 3))
}

func minimizeMax(nums []int, p int) int {
	if len(nums) < 2 || p == 0 {
		return 0
	}

	result := math.MaxInt
	sort.Ints(nums)

	leftDiffBound := 0
	rightDiffBound := nums[len(nums)-1] - nums[0]
	var currentGuess int
	for leftDiffBound <= rightDiffBound {
		currentGuess = (leftDiffBound + rightDiffBound) / 2
		diffCount := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= currentGuess {
				diffCount++
				i++
			}
		}
		if diffCount >= p {
			if currentGuess < result {
				result = currentGuess
			}
			rightDiffBound = currentGuess - 1
		} else {
			leftDiffBound = currentGuess + 1
		}
	}

	return result
}

func minimizeMaxDP(nums []int, p int) int {
	if len(nums) < 2 {
		return 0
	}
	if p == 0 {
		return 0
	}

	sort.Ints(nums)

	result := 0
	for i := 0; i < p; i++ {
		differences := make([][]int, 2)
		currentMin := math.MaxInt
		minIndex := 0

		// Index row
		differences[0] = make([]int, len(nums)-1)
		// Value row
		differences[1] = make([]int, len(nums)-1)

		// Filling rows
		for j := 0; j < len(differences[0]); j++ {
			differences[0][j] = j
			differences[1][j] = nums[j+1] - nums[j]
			if differences[1][j] < currentMin {
				currentMin = differences[1][j]
				minIndex = j
			}
		}
		if differences[1][minIndex] > result {
			result = differences[1][minIndex]
		}
		// Removing those slice members
		nums = append(nums[0:minIndex], nums[minIndex+2:]...)
	}
	return result
}
