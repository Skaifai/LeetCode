package main

func main() {
	costs := []int{4, 9, 3, 8, 8, 9}
	rope := "bbbaaa"
	print(minCost(rope, costs))
}

func minCost(colors string, neededTime []int) int {
	result := 0
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			sum := 0
			middleCost := 0
			leftCost := neededTime[i]
			rightCost := neededTime[i+1]
			for i < len(colors)-1 && colors[i] == colors[i+1] {
				sum += neededTime[i]
				if neededTime[i] > middleCost {
					middleCost = neededTime[i]
				}
				i++
			}
			if i < len(colors) {
				sum += neededTime[i]
			}
			rightCost = neededTime[i]
			minimum := max(leftCost, max(middleCost, rightCost))
			result += sum - minimum
		}
	}
	return result
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
