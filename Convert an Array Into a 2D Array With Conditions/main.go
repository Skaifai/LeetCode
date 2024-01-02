package main

import "fmt"

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}

func findMatrix(nums []int) [][]int {
	var result [][]int
	integers := make(map[int]int)

	for _, num := range nums {
		if len(result) == integers[num] {
			result = append(result, []int{})
		}
		result[integers[num]] = append(result[integers[num]], num)
		integers[num]++
	}

	return result
}

//func findMatrix(nums []int) [][]int {
//	var result [][]int
//	integers := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		integers[nums[i]]++
//	}
//
//	for integer := range integers {
//		for i := 0; i < len(result) && i < integers[integer]; i++ {
//			result[i] = append(result[i], integer)
//		}
//
//		for len(result) < integers[integer] {
//			result = append(result, []int{integer})
//		}
//	}
//
//	return result
//}

//func findMatrix(nums []int) [][]int {
//	var result [][]int
//	count := 0
//	integers := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		integers[nums[i]]++
//		count++
//	}
//
//	for count > 0 {
//		intermediate := make([]int, 0)
//		for integer := range integers {
//			if integers[integer] > 0 {
//				count--
//				intermediate = append(intermediate, integer)
//				integers[integer]--
//			}
//		}
//		result = append(result, intermediate)
//	}
//	return result
//}
