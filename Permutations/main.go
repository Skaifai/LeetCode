package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(permute(slice))
}

func permute(nums []int) [][]int {
	var result [][]int
	result = append(result, nums)
	if len(nums) == 1 {
		return result
	}

	checkLength := 1
	for i := 2; i <= len(nums); i++ {
		checkLength *= i
	}

	result = recursion(nums)

	return result
}

func recursion(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	checkLength := 1
	for i := 2; i <= len(nums); i++ {
		checkLength *= i
	}

	result := make([][]int, 0)

	var generatedSlices [][]int
	for i := 0; i < len(nums); i++ {
		copyOfNums := make([]int, len(nums))
		copy(copyOfNums, nums)
		slicesToAdd := recursion(append(copyOfNums[0:i], copyOfNums[i+1:]...))
		for j := 0; j < len(slicesToAdd); j++ {
			arrayToAdd := append([]int{nums[i]}, slicesToAdd[j]...)
			generatedSlices = append(generatedSlices, arrayToAdd)
		}
	}
	result = append(result, generatedSlices...)
	//temp := nums[0 : len(nums)-1]
	//generatedSlices = append(generatedSlices, recursion(temp)...)

	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < checkLength; j++ {
	//		result[i] = append([]int{nums[i]}, generatedSlices[j]...)
	//	}
	//}

	return result
}
