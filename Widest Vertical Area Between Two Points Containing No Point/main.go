package main

import "sort"

func main() {

}

//func maxWidthOfVerticalArea(points [][]int) int {
//	flattenedArray := make([]int, len(points))
//	for i := 0; i < len(points); i++ {
//		flattenedArray[i] = points[i][0]
//	}
//
//	sort.Ints(flattenedArray)
//
//	difference := 0
//	for i := 1; i < len(flattenedArray); i++ {
//		if flattenedArray[i]-flattenedArray[i-1] > difference {
//			difference = flattenedArray[i] - flattenedArray[i-1]
//		}
//	}
//
//	return difference
//}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	difference := 0
	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i-1][0] > difference {
			difference = points[i][0] - points[i-1][0]
		}
	}

	return difference
}
