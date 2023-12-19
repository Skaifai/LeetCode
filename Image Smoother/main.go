package main

import "fmt"

func main() {
	example := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	fmt.Println(imageSmoother(example))
}

func imageSmoother(img [][]int) [][]int {
	result := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		result[i] = make([]int, len(img[0]))
	}

	for currentRow := 0; currentRow < len(img); currentRow++ {
		for currentColumn := 0; currentColumn < len(img[0]); currentColumn++ {
			sum := 0
			numConsidered := 0
			for k := -1; k < 2; k++ {
				if currentRow+k < 0 {
					continue
				} else if currentRow+k >= len(img) {
					continue
				}
				for l := -1; l < 2; l++ {
					if currentColumn+l < 0 {
						continue
					} else if currentColumn+l >= len(img[0]) {
						continue
					}
					sum += img[currentRow+k][currentColumn+l]
					numConsidered++
				}
			}
			if numConsidered == 0 {
				continue
			}
			average := sum / numConsidered
			result[currentRow][currentColumn] = average
		}
	}
	return result
}
