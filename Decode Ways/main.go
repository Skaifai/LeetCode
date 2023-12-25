package main

import "strconv"

func main() {
	print(numDecodings("2101"))
}

func numDecodings(s string) int {
	counter1 := 0
	counter2 := 0
	previousCounter2 := counter2
	if s[0] != '0' {
		counter1 = 1
	}
	for i := 1; i < len(s); i++ {
		previousCounter2 = counter2
		if s[i-1] != '0' {
			number, _ := strconv.Atoi(s[i-1 : i+1])
			if number <= 26 {
				counter2 = counter1
			} else {
				counter2 = 0
			}
		} else {
			counter2 = 0
		}

		if s[i] != '0' {
			counter1 = previousCounter2 + counter1
		} else {
			counter1 = 0
		}
	}

	return counter1 + counter2
}
