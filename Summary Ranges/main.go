package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	r := make([]string, 0)

	if len(nums) == 1 {
		r = append(r, strconv.Itoa(nums[0]))
		return r
	}

	currentIndex := 0
	for currentIndex < len(nums) {
		currentValue := nums[currentIndex]
		stringToAdd := strconv.Itoa(currentValue)

		if len(nums) <= currentIndex+1 {
			r = append(r, stringToAdd)
			return r
		}

		if nums[currentIndex+1] != currentValue+1 {
			r = append(r, stringToAdd)
			currentIndex++
			continue
		}

		for currentIndex < len(nums)-1 {
			if nums[currentIndex+1] == currentValue+1 {
				currentIndex++
				currentValue++
				continue
			}
			break
		}
		stringToAdd = stringToAdd + "->" + strconv.Itoa(currentValue)

		r = append(r, stringToAdd)

		currentIndex++
	}

	return r
}
