package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}))
}

func minOperations(nums []int) int {
	result := 0

	occurrences := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		occurrences[nums[i]]++
	}

	for _, number := range occurrences {
		if number == 1 {
			return -1
		}
		var intermediate int

		if number%3 == 2 || number%3 == 1 {
			intermediate++
		}

		intermediate += number / 3

		result += intermediate
	}

	return result
}
