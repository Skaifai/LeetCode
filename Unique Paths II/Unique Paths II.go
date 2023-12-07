package main

import "fmt"

func main() {
	test := [][]int{{0, 0}, {1, 1}, {0, 0}}
	fmt.Println(uniquePathsWithObstacles(test))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dpTable := make([][]int, m)
	for i := 0; i < m; i++ {
		dpTable[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dpTable[i][j] = -1
			}
		}
	}

	dpTable[0][0] = 1

	for i := 1; i < len(dpTable[0]); i++ {
		if dpTable[0][i] == -1 {
			dpTable[0][i] = 0
		} else {
			dpTable[0][i] += dpTable[0][i-1]
		}
	}
	for i := 1; i < len(dpTable); i++ {
		if dpTable[i][0] == -1 {
			dpTable[i][0] = 0
		} else {
			dpTable[i][0] += dpTable[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dpTable[i][j] == -1 {
				dpTable[i][j] = 0
			} else {
				dpTable[i][j] = dpTable[i-1][j] + dpTable[i][j-1]
			}
		}
	}

	return dpTable[m-1][n-1]
}
