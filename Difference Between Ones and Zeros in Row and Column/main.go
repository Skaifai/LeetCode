package main

func main() {
	slice := [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}
	print(onesMinusZeros(slice))
}

func onesMinusZeros(grid [][]int) [][]int {
	result := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[0]))
	}

	onesRow := make([]int, len(grid))
	zerosRow := make([]int, len(grid))

	onesColumn := make([]int, len(grid[0]))
	zerosColumn := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				onesRow[i]++
				onesColumn[j]++
			}
		}
	}

	for i := 0; i < len(onesRow); i++ {
		zerosRow[i] = len(onesRow) - onesRow[i]
	}

	for i := 0; i < len(onesColumn); i++ {
		zerosColumn[i] = len(onesColumn) - onesColumn[i]
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// diff[i][j] = onesRowi + onesColj - zerosRowi - zerosColj
			result[i][j] = onesRow[i] + onesColumn[j] - zerosRow[i] - zerosColumn[j]
		}
	}

	return result
}
