package main

import "fmt"

func main() {
	x := 12343211
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverseNumber := 0
	copyOfX := x
	for copyOfX > 0 {
		digit := copyOfX % 10
		copyOfX /= 10
		reverseNumber *= 10
		reverseNumber += digit
	}

	if reverseNumber == x {
		return true
	}

	return false
}
