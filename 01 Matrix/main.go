package main

import "fmt"

func main() {
	test := [][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}
	fmt.Println(updateMatrix(test))
}

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[0]))
	}

	// First pass. Left to right, top to bottom
	// First line
	addNumber := 1
	for i := 1; i < len(mat[0]); i++ {
		if mat[0][i] != 0 {
			res[0][i] = addNumber
			addNumber++
		} else {
			res[0][i] = 0
			addNumber = 1
		}
	}
	// First column
	addNumber = 1
	for i := 1; i < len(mat); i++ {
		if mat[i][0] != 0 {
			res[i][0] = addNumber
			addNumber++
		} else {
			res[i][0] = 0
			addNumber = 1
		}
	}
	// Everything else
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				if res[i-1][j] < res[i][j-1] {
					res[i][j] = res[i-1][j] + 1
				} else {
					res[i][j] = res[i][j-1] + 1
				}
			}
		}
	}

	// Second pass
	// Last line
	addNumber = res[len(res)-1][len(res[0])-1]
	for i := len(mat[0]) - 1; i >= 0; i-- {
		if mat[len(mat)-1][i] != 0 {
			if addNumber < res[len(res)-1][i] {
				res[len(res)-1][i] = addNumber
			}
			addNumber++
		} else {
			res[len(res)-1][i] = 0
			addNumber = 1
		}
	}
	// Last column
	addNumber = res[len(res)-1][len(res[0])-1]
	for i := len(mat) - 1; i >= 0; i-- {
		if mat[i][len(mat[0])-1] != 0 {
			if addNumber < res[i][len(res[0])-1] {
				res[i][len(res)-1] = addNumber
			}
			addNumber++
		} else {
			res[i][len(res[0])-1] = 0
			addNumber = 1
		}
	}

	// Everything else
	for i := len(mat) - 2; i >= 0; i-- {
		for j := len(mat[0]) - 2; j >= 0; j-- {
			potential := 0
			if res[i][j] != 0 {
				if res[i+1][j] < res[i][j+1] {
					potential = res[i+1][j] + 1
				} else {
					potential = res[i][j+1] + 1
				}
				if potential < res[i][j] {
					res[i][j] = potential
				}
			}
		}
	}

	if len(res) > 1 && len(res[0]) > 1 && mat[0][0] != 0 {
		if res[0][1] < res[1][0] {
			res[0][0] = res[0][1] + 1
		} else {
			res[0][0] = res[1][0] + 1
		}
	}

	return res
}
