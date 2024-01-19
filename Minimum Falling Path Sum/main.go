package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(arr))
}

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	minSum := math.MaxInt

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			min := math.MaxInt
			for k := -1; k <= 1; k++ {
				checkColumn := j + k
				if checkColumn < 0 || checkColumn > len(dp[j]) {
					continue
				}
				if dp[i-1][checkColumn] < min {
					min = dp[i-1][checkColumn]
				}
			}
			dp[i][j] = min + matrix[i][j]
		}
	}

	for i := 0; i < len(dp[len(dp)-1]); i++ {
		if dp[len(dp)-1][i] < minSum {
			minSum = dp[len(dp)-1][i]
		}
	}

	return minSum
}
