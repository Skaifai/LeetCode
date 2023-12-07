package main

func main() {

}

func climbStairsRecursive(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2
	return climbStairsMemo(n, memo)
}

func climbStairsMemo(n int, memo map[int]int) int {
	val, ok := memo[n]
	if ok {
		return val
	}
	result := climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	memo[n] = result
	return result
}
