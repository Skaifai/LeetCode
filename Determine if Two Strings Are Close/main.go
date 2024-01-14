package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(closeStrings("abc", "bca"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1map := make(map[uint8]int)
	word2map := make(map[uint8]int)

	for i := 0; i < len(word1); i++ {
		word1map[word1[i]]++
		word2map[word2[i]]++
	}
	if len(word1map) != len(word2map) {
		return false
	}

	sliceOfNums1 := make([]int, len(word1map))
	sliceOfLetters1 := make([]uint8, len(word1map))
	counter := 0
	for u := range word1map {
		sliceOfLetters1[counter] = u
		sliceOfNums1[counter] = word1map[u]
		counter++
	}
	sort.Ints(sliceOfNums1)
	sort.Slice(sliceOfLetters1, func(i, j int) bool {
		return sliceOfLetters1[j] > sliceOfLetters1[i]
	})
	sliceOfNums2 := make([]int, len(word2map))
	sliceOfLetters2 := make([]uint8, len(word1map))
	counter = 0
	for u := range word2map {
		sliceOfLetters2[counter] = u
		sliceOfNums2[counter] = word2map[u]
		counter++
	}
	sort.Ints(sliceOfNums2)
	sort.Slice(sliceOfLetters2, func(i, j int) bool {
		return sliceOfLetters2[j] > sliceOfLetters2[i]
	})
	for i := 0; i < len(sliceOfNums1); i++ {
		if sliceOfLetters1[i] != sliceOfLetters2[i] {
			return false
		}
		if sliceOfNums1[i] != sliceOfNums2[i] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
