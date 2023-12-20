package main

import "math"

func main() {
	array := []int{1, 2, 2}
	print(buyChoco(array, 3))
}

func buyChoco(prices []int, money int) int {
	leftover := money

	firstSmallest := math.MaxInt
	secondSmallest := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < firstSmallest {
			secondSmallest = firstSmallest
			firstSmallest = prices[i]
		} else if prices[i] < secondSmallest {
			secondSmallest = prices[i]
		}
	}

	leftover -= firstSmallest + secondSmallest

	if leftover < 0 {
		return money
	}

	return leftover
}
