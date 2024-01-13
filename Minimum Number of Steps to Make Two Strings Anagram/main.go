package main

import "fmt"

func main() {
	fmt.Println(minSteps("abb", "aab"))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minSteps(s string, t string) int {
	res := 0

	word := map[uint8]int{}

	for i := 0; i < len(s); i++ {
		word[s[i]-'a']++
		word[t[i]-'a']--
	}

	for _, i := range word {
		res += abs(i)
	}

	return res / 2
}

//func minSteps(s string, t string) int {
//	res := 0
//
//	var word [26]int
//
//	for i := 0; i < len(s); i++ {
//		word[s[i]-'a']++
//		word[t[i]-'a']--
//	}
//
//	for i := 0; i < 26; i++ {
//		res += abs(word[i])
//	}
//
//	return res / 2
//}
