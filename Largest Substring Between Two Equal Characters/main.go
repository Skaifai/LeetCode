package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("mgntdygtxrvxjnwksqhxuxtrv"))
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxLengthBetweenEqualCharacters(s string) int {
	firstSeen := make(map[rune]int)
	result := -1

	for i := len(s) - 1; i >= 0; i-- {
		letter := rune(s[i])
		firstSeen[letter] = i
	}

	for i, letter := range s {
		intermediateResult := i - firstSeen[letter] - 1
		result = max(intermediateResult, result)
	}

	return result
}

//func maxLengthBetweenEqualCharacters(s string) int {
//	firstSeen := make(map[rune]int)
//	lastSeen := make(map[rune]int)
//	result := -1
//
//	for i := len(s) - 1; i >= 0; i-- {
//		letter := rune(s[i])
//		firstSeen[letter] = i
//	}
//
//	for letter, _ := range firstSeen {
//		intermediateResult := lastSeen[letter] - firstSeen[letter] - 1
//		result = max(result, intermediateResult)
//	}
//
//	return result
//}

//func maxLengthBetweenEqualCharacters(s string) int {
//	occurrence := make(map[rune]int)
//	lastSeen := make(map[rune]int)
//	result := -1
//
//	for _, letter := range s {
//		occurrence[letter]++
//	}
//
//	for i, letter := range s {
//		if occurrence[letter] == 1 {
//			continue
//		}
//		intermediateResult := i - lastSeen[letter] - 1
//		result = max(result, intermediateResult)
//		lastSeen[letter] = i
//	}
//
//	return result
//}
