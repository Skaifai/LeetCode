package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
	//fmt.Println(numberOfArithmeticSlices([]int{1, 1, 1, 1}))
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	res := 0

	dp := make([][]int, len(nums)-1)
	differences := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i] = append(dp[i], nums[j]-nums[i])
			differences[nums[j]-nums[i]]++
		}
	}

	fmt.Println(dp)
	fmt.Println(differences)

	for key, value := range differences {
		if value >= 2 {
			count := 0
			for i := 0; i < len(dp); i++ {
				for j := 0; j < len(dp[i]); j++ {
					if key == dp[i][j] {
						count++
						i = j
						j = -1
					}
				}
			}
			res += count
		}
	}
	return res
}
