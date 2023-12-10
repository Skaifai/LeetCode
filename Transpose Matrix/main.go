package Transpose_Matrix

func main() {

}

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	columns := len(matrix[0])

	newMatrix := make([][]int, columns)

	for i := 0; i < columns; i++ {
		newMatrix[i] = make([]int, rows)
	}

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}

	return newMatrix
}
