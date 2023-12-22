package main

func main() {
	print(maxScore("011101"))
	print(maxScore("0111011"))
	print(maxScore("00011111"))
	print(maxScore("111110000"))
}

func maxScore(s string) int {
	result := 0

	for i := 1; i < len(s); i++ {
		currentScore := 0
		for j := 0; j < i; j++ {
			if s[j] == '0' {
				currentScore++
			}
		}
		for j := i; j < len(s); j++ {
			if s[j] == '1' {
				currentScore++
			}
		}
		if currentScore > result {
			result = currentScore
		}
	}

	return result
}
