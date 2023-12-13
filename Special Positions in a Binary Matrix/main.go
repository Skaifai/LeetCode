package main

func main() {
	//slice := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	slice := [][]int{{0, 0, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 1, 0, 0}}
	print(numSpecial(slice))
}

func numSpecial(mat [][]int) int {
	counter := 0
	horizontalSums := make([]int, len(mat))
	verticalSums := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		currentSum := 0
		for j := 0; j < len(mat[0]); j++ {
			currentSum += mat[i][j]
		}
		horizontalSums[i] = currentSum
	}
	for i := 0; i < len(mat[0]); i++ {
		currentSum := 0
		for j := 0; j < len(mat); j++ {
			currentSum += mat[j][i]
		}
		verticalSums[i] = currentSum
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && horizontalSums[i] == 1 && verticalSums[j] == 1 {
				counter++
			}
		}
	}
	return counter
}

//func numSpecial(mat [][]int) int {
//	counter := 0
//	for i := 0; i < len(mat); i++ {
//		for j := 0; j < len(mat[0]); j++ {
//			hor := j
//			for j < len(mat[0]) && mat[i][j] == 0 {
//				j++
//			}
//			j++
//			if j < len(mat[0]) {
//				for j < len(mat[0]) && mat[i][j] == 0 {
//					j++
//				}
//				if j < len(mat[0]) {
//					j = len(mat[0])
//					break
//				} else {
//					j = hor
//					for k := 0; k < len(mat); k++ {
//						for k < len(mat) && mat[k][j] == 0 {
//							k++
//						}
//						k++
//						if k < len(mat) {
//							for k < len(mat) && mat[k][j] == 0 {
//								k++
//							}
//							if k < len(mat) {
//								k = len(mat)
//								break
//							} else {
//								counter++
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//	return counter
//}
