package main

import (
	"fmt"
)

func main() {
	slice := []byte{'a', 'a', 'a', 'b', 'a', 'a', 'a'}
	fmt.Println(compress(slice))
	fmt.Println(slice)
}

func compress(chars []byte) int {
	index := 0
	for i := 0; i < len(chars); i++ {
		currentLetter := chars[i]
		currentLetterCount := 1
		for i < len(chars)-1 && chars[i+1] == chars[i] {
			currentLetterCount++
			i++
		}
		if currentLetterCount > 1 {
			chars[index] = currentLetter
			index++

			//Count length of the number
			lengthOfNumber := 0
			number := currentLetterCount
			for number != 0 {
				number /= 10
				lengthOfNumber++
			}
			index += lengthOfNumber - 1
			number = currentLetterCount
			for number != 0 {
				chars[index] = '1' + byte(number%10) - 1
				index--
				number /= 10
			}
			index += lengthOfNumber + 1
		} else {
			chars[index] = currentLetter
			index++
		}
	}

	return index
}
