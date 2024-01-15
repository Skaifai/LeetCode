package main

import "fmt"

func main() {
	test := "DCXXI"
	fmt.Println(romanToInt(test))
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 >= len(s) {
				res++
				break
			}
			if s[i+1] == 'V' {
				i++
				res += 4
			} else if s[i+1] == 'X' {
				i++
				res += 9
			} else {
				res++
			}
		case 'V':
			res += 5
		case 'X':
			if i+1 >= len(s) {
				res += 10
				break
			}
			if s[i+1] == 'L' {
				i++
				res += 40
			} else if s[i+1] == 'C' {
				i++
				res += 90
			} else {
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			if i+1 >= len(s) {
				res += 100
				break
			}
			if s[i+1] == 'D' {
				i++
				res += 400
			} else if s[i+1] == 'M' {
				i++
				res += 900
			} else {
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}

	}
	return res
}

//
//func romanToInt(s string) int {
//	result := 0
//
//	charsMap := make(map[rune]int)
//	charsMap['V'] = 5
//	charsMap['L'] = 50
//	charsMap['D'] = 500
//	charsMap['M'] = 1000
//
//	sliceOfCharacters := make([]rune, 0)
//	for _, letter := range s {
//		sliceOfCharacters = append(sliceOfCharacters, letter)
//	}
//
//	for i := 0; i < len(sliceOfCharacters); i++ {
//		if sliceOfCharacters[i] == 'I' {
//			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'V' {
//				result += 4
//				i++
//			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'X' {
//				result += 9
//				i++
//			} else {
//				result += 1
//			}
//		} else if sliceOfCharacters[i] == 'X' {
//			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'L' {
//				result += 40
//				i++
//			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'C' {
//				result += 90
//				i++
//			} else {
//				result += 10
//			}
//		} else if sliceOfCharacters[i] == 'C' {
//			if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'D' {
//				result += 400
//				i++
//			} else if i < len(sliceOfCharacters)-1 && sliceOfCharacters[i+1] == 'M' {
//				result += 900
//				i++
//			} else {
//				result += 100
//			}
//		} else {
//			result += charsMap[sliceOfCharacters[i]]
//		}
//	}
//
//	return result
//}
