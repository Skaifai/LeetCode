package main

func main() {
	totalMoney(10)
}

func totalMoney(n int) int {
	lastDeposit := 0
	result := 0
	if n/7 > 0 {
		lastDeposit = n / 7
		for i := 0; i < n/7; i++ {
			result += 28 + i*7
		}
	}

	if n%7 > 0 {
		for i := 0; i < n%7; i++ {
			result += lastDeposit + 1 + i
		}
	}

	return result
}
