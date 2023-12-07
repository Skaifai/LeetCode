package main

import "fmt"

func main() {
	test := "hello"
	note := "helloq"
	fmt.Println(canConstruct(note, test))
}

func canConstruct(ransomNote string, magazine string) bool {
	// Fool-proofing, check if magazine has more letters than ransomNote
	if len(magazine) < len(ransomNote) {
		return false
	}

	lettersCount := make(map[rune]int, 26)
	for _, ch := range magazine {
		lettersCount[ch]++
	}
	for _, ch := range ransomNote {
		lettersCount[ch]--
		if lettersCount[ch] < 0 {
			return false
		}
	}
	return true
}
