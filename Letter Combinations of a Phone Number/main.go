package main

import "fmt"

func main() {
	test := ""
	fmt.Println(letterCombinations(test))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	result := make([]string, 0)

	dictionary := make(map[int]string)
	dictionary[2] = "abc"
	dictionary[3] = "def"
	dictionary[4] = "ghi"
	dictionary[5] = "jkl"
	dictionary[6] = "mno"
	dictionary[7] = "pqrs"
	dictionary[8] = "tuv"
	dictionary[9] = "wxyz"

	totalRepetitions := 1
	for _, digit := range digits {
		digitInt := int(digit - '0')
		totalRepetitions *= len(dictionary[digitInt])
	}

	repetitionsLeft := totalRepetitions
	for i := 0; i < totalRepetitions; i++ {
		result = append(result, "")
	}

	for _, digit := range digits {
		digitInt := int(digit - '0')
		repetitionsLeft = repetitionsLeft / len(dictionary[digitInt])
		currentLetterIndex := 0
		for i := 0; i < totalRepetitions; {
			if currentLetterIndex >= len(dictionary[digitInt]) {
				currentLetterIndex = 0
			}
			for j := 0; j < repetitionsLeft; j++ {
				word := dictionary[digitInt]
				result[i] = result[i] + string(word[currentLetterIndex])
				i++
			}
			currentLetterIndex++
		}
	}

	return result
}
