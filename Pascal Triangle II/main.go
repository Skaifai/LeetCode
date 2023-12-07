package main

func main() {

}

func getRow(rowIndex int) []int {
	previousRow := []int{1}
	currentRow := []int{1, 1}
	if rowIndex == 0 {
		return previousRow
	} else {
		for i := 0; i <= rowIndex; i++ {
			previousRow = currentRow
			currentRow = make([]int, i+1)
			currentRow[0] = 1
			currentRow[len(currentRow)-1] = 1
			for j := 1; j < len(currentRow)-1; j++ {
				currentRow[j] = previousRow[j-1] + previousRow[j]
			}
		}
		return currentRow
	}
}
