package main

func main() {

}

func maximumWealth(accounts [][]int) int {
	r := 0

	for i := 0; i < len(accounts); i++ {
		currentSum := 0
		for j := 0; j < len(accounts[i]); j++ {
			currentSum += accounts[i][j]
		}

		if currentSum > r {
			r = currentSum
		}
	}

	return r
}
