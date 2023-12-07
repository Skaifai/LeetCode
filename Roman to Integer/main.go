package main

import "fmt"

func main() {
	test := "DCXXI"
	fmt.Println(romanToInt(test))
}

func romanToInt(s string) int {
	result := 0

	charsMap := make(map[rune]int)
	charsMap['V'] = 5
	charsMap['L'] = 50
	charsMap['D'] = 500
	charsMap['M'] = 1000

	sliceOfCharacters := make([]rune, 0)
	for _, letter := range s {
		sliceOfCharacters = append(sliceOfCharacters, letter)
	}

	for i := 0; i < len(sliceOfCharacters); i++ {
		if sliceOfCharacters[i] == 'I' {
			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'V' {
				result += 4
				i++
			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'X' {
				result += 9
				i++
			} else {
				result += 1
			}
		} else if sliceOfCharacters[i] == 'X' {
			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'L' {
				result += 40
				i++
			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'C' {
				result += 90
				i++
			} else {
				result += 10
			}
		} else if sliceOfCharacters[i] == 'C' {
			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'D' {
				result += 400
				i++
			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'M' {
				result += 900
				i++
			} else {
				result += 100
			}
		} else {
			result += charsMap[sliceOfCharacters[i]]
		}
	}

	return result
}
