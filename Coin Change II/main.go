package main

import (
	"fmt"
	"sort"
)

func main() {
	test2 := []int{3, 5, 7, 8, 9, 10, 11, 501, 502}
	fmt.Println(change(500, test2))
}

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	} else if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	// Check which coins are even allowed
	i := len(coins) - 1

	for i >= 0 && coins[i] > amount {
		i--
	}
	if i < 0 {
		return 0
	}

	result := 0
	result += change(amount-coins[i], coins[:i+1])

	result += change(amount, coins[:i])
	return result
}

func changeDP(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	result := 0
	sort.Ints(coins)
	tableOfMaxAllowedUses := make([]int, 0)
	allowedCoins := make([]int, 0)
	maxPossibleAnswer := 1
	for i := 0; i < len(coins); i++ {
		maxAllowedUses := amount / coins[i]
		if maxAllowedUses > 0 {
			tableOfMaxAllowedUses = append(tableOfMaxAllowedUses, maxAllowedUses)
			allowedCoins = append(allowedCoins, coins[i])
			maxPossibleAnswer *= tableOfMaxAllowedUses[i] + 1
		}
	}
	maxPossibleAnswer /= tableOfMaxAllowedUses[0] + 1

	multipliersTable := make([]int, len(allowedCoins))
	for i := 0; i < maxPossibleAnswer; i++ {
		for j := 0; j < tableOfMaxAllowedUses[0]+1; j++ {
			currentSum := 0

			// Find sum
			for k := 0; k < len(multipliersTable); k++ {
				currentSum += allowedCoins[k] * multipliersTable[k]
			}

			if currentSum >= amount {
				if currentSum == amount {
					result++
				}
				for j < tableOfMaxAllowedUses[0]+1 {
					updateArrayCounter(multipliersTable, tableOfMaxAllowedUses)
					j++
				}
				break
			}

			updateArrayCounter(multipliersTable, tableOfMaxAllowedUses)

		}

	}

	return result
}

func updateArrayCounter(arr []int, maxValues []int) {
	arr[0]++
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > maxValues[i] {
			arr[i+1]++
			arr[i] = 0
		}
	}
	if arr[len(arr)-1] > maxValues[len(arr)-1] {
		for i := 0; i < len(arr); i++ {
			arr[i] = 0
		}
	}
}
