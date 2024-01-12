package main

import "strings"

func main() {

}

func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2:]

	count := 0
	for i := 0; i < len(firstHalf); i++ {
		if strings.Contains(vowels, string(firstHalf[i])) {
			count++
		}
	}

	for i := 0; i < len(secondHalf); i++ {
		if strings.Contains(vowels, string(secondHalf[i])) {
			count--
		}
	}

	return count == 0
}
