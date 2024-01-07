package main

import (
	"fmt"
	"strconv"
)

func main() {
	getLengthOfOptimalCompression("ababa", 3)
}

func getLengthOfOptimalCompression(s string, k int) int {
	if k == len(s) {
		return 0
	}
	compressed, letters, _ := compressString(s)
	if k == 0 {
		return len(compressed)
	}

	var prefixSums = map[rune][]int{}
	currentIndex := 0
	for _, letter := range letters {
		for len(prefixSums[letter]) < len(s) {
			prefixSums[letter] = append(prefixSums[letter], 0)
		}
	}
	for i := 0; i < len(s); i++ {
		prefixSums[rune(s[i])][i] = 1
	}

	for _, slice := range prefixSums {
		runningSum := 0
		for i := 0; i < len(slice); i++ {
			runningSum += slice[i]
			slice[i] = runningSum
		}
	}

	fmt.Println(prefixSums)

	return 0
}

func compressString(s string) (string, []rune, []int) {
	compressedString := ""
	var letters []rune
	var quantity []int

	for i := 0; i < len(s); i++ {
		currentLetter := s[i]
		currentLetterCount := 1
		for i < len(s)-1 && s[i+1] == s[i] {
			currentLetterCount++
			i++
		}
		if currentLetterCount > 1 {
			compressedString = compressedString + string(currentLetter)
			compressedString = compressedString + strconv.Itoa(currentLetterCount)
			letters = append(letters, rune(currentLetter))
			quantity = append(quantity, currentLetterCount)
		} else if i == len(s) {
			compressedString = compressedString + string(currentLetter)
			compressedString = compressedString + strconv.Itoa(currentLetterCount)
			letters = append(letters, rune(currentLetter))
			quantity = append(quantity, currentLetterCount)
		} else {
			compressedString = compressedString + string(currentLetter)
			letters = append(letters, rune(currentLetter))
			quantity = append(quantity, currentLetterCount)
		}
	}

	return compressedString, letters, quantity
}
