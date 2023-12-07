package main

import "strconv"

func main() {

}

func fizzBuzz(n int) []string {
	r := make([]string, n)

	for i := 1; i <= n; i++ {

		currentString := ""

		if i%3 == 0 {
			currentString = currentString + "Fizz"
		}

		if i%5 == 0 {
			currentString = currentString + "Buzz"
		}

		if i%3 != 0 && i%5 != 0 {
			currentString = currentString + strconv.Itoa(i)
		}

		r[i-1] = currentString
	}

	return r
}
