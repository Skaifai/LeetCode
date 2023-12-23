package main

func main() {
	print(isPathCrossing("NESW"))
}

type yCoordinates struct {
	//TODO
}

func isPathCrossing(path string) bool {
	xCord := 0
	yCord := 0
	visited := make(map[int]map[int]int)
	visited[0] = make(map[int]int)
	visited[0][0] = 1

	for i := 0; i < len(path); i++ {
		if path[i] == 'N' {
			yCord++
		} else if path[i] == 'S' {
			yCord--
		} else if path[i] == 'E' {
			xCord++
		} else {
			xCord--
		}

		if visited[xCord] == nil {
			visited[xCord] = make(map[int]int)
			visited[xCord][yCord]++
			continue
		} else {
			if visited[xCord][yCord] == 1 {
				return true
			} else {
				visited[xCord][yCord]++
			}
		}
	}

	return false
}

//func isPathCrossing(path string) bool {
//	counter := 0
//	for i := 1; i < len(path); i++ {
//		if path[i-1] == 'N' {
//			if path[i] == 'S' {
//				return true
//			} else if path[i] == 'E' {
//				counter++
//			} else if path[i] == 'W' {
//				counter--
//			} else if path[i] == 'N' {
//				continue
//			}
//		} else if path[i-1] == 'S' {
//			if path[i] == 'S' {
//				continue
//			} else if path[i] == 'E' {
//				counter--
//			} else if path[i] == 'W' {
//				counter++
//			} else if path[i] == 'N' {
//				return true
//			}
//		} else if path[i-1] == 'W' {
//			if path[i] == 'S' {
//				counter--
//			} else if path[i] == 'E' {
//				return true
//			} else if path[i] == 'W' {
//				continue
//			} else if path[i] == 'N' {
//				counter++
//			}
//		} else if path[i-1] == 'E' {
//			if path[i] == 'S' {
//				counter++
//			} else if path[i] == 'E' {
//				continue
//			} else if path[i] == 'W' {
//				return true
//			} else if path[i] == 'N' {
//				counter--
//			}
//		}
//
//		if counter >= 3 || counter <= -3 {
//			return true
//		} else if counter == 0 {
//			i--
//		}
//	}
//	return false
//}
