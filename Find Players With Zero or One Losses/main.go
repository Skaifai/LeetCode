package main

import (
	"fmt"
	"sort"
)

func main() {
	slices := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	fmt.Println(findWinners(slices))

}

func findWinners(matches [][]int) [][]int {
	wins := make(map[int]int)
	losses := make(map[int]int)

	for i := 0; i < len(matches); i++ {
		wins[matches[i][0]]++
		losses[matches[i][1]]++
	}

	zeroLosses := make([]int, 0)
	oneLoss := make([]int, 0)

	for i, num := range losses {
		if num == 1 {
			oneLoss = append(oneLoss, i)
		}
	}

	for i := range wins {
		if losses[i] == 0 {
			zeroLosses = append(zeroLosses, i)
		}
	}

	sort.Ints(zeroLosses)
	sort.Ints(oneLoss)

	return [][]int{zeroLosses, oneLoss}
}

//func findWinners(matches [][]int) [][]int {
//	zeroLosses := make([]int, 0)
//	oneLoss := make([]int, 0)
//	losers := make([]int, 0)
//
//	for i := 0; i < len(matches); i++ {
//		winner := matches[i][0]
//		loser := matches[i][1]
//		if sort.SearchInts(losers, winner) < len(losers) && sort.SearchInts(losers, loser) < len(losers) {
//			continue
//		}
//
//		// If loser already has one loss, remove him to the loser bracket
//		cache := sort.SearchInts(oneLoss, loser)
//		if cache < len(losers) && losers[cache] != loser {
//			oneLoss = append(oneLoss[:cache], oneLoss[cache+1:]...)
//			losers = insertPlayer(loser, losers)
//		} else {
//			// If loser was in winners, then remove him from the winners and move to oneLoss
//			cache = sort.SearchInts(zeroLosses, loser)
//			if cache < len(zeroLosses) && zeroLosses[cache] != loser {
//				zeroLosses = append(zeroLosses[:cache], zeroLosses[cache+1:]...)
//				oneLoss = insertPlayer(loser, losers)
//			}
//		}
//		// Check if winner is in loser, if yes, then continue
//		cache = sort.SearchInts(losers, winner)
//		if cache < len(losers) && losers[cache] == winner {
//			continue
//		}
//
//		// Check if winner is already in winners, if not, then insert
//		cache = sort.SearchInts(zeroLosses, winner)
//		if cache < len(zeroLosses) && losers[cache] != winner {
//			zeroLosses = insertPlayer(winner, zeroLosses)
//			continue
//		}
//
//		if cache == len(zeroLosses) {
//			zeroLosses = append(zeroLosses, winner)
//		}
//	}
//
//	return [][]int{zeroLosses, oneLoss}
//}

func insertPlayer(player int, slice []int) []int {
	insertionIndex := sort.SearchInts(slice, player)

	for i := insertionIndex; i < len(slice); i++ {
		prev := slice[i]
		slice[i] = player
		player = prev
	}
	slice = append(slice, player)
	return slice
}
