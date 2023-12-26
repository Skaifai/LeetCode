package main

func main() {
	print(numRollsToTarget(3, 6, 6))
}

func numRollsToTarget(n int, k int, target int) int {

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, target)
		for j := 0; j < target; j++ {
			memo[i][j] = -1
		}
	}
	result := recursiveWithTable(n, k, target, memo)

	//fmt.Println(memo)

	return result
}

func recursiveWithTable(n int, k int, target int, table [][]int) int {
	if target == 0 {
		return 0
	}

	if n == 0 {
		return 0
	}

	if target > n*k || n > target {
		table[n-1][target-1] = 0
		return 0
	}

	if n == 1 {
		table[n-1][target-1] = 1
		return 1
	}

	if table[n-1][target-1] != -1 {
		return table[n-1][target-1]
	}

	const modulo = 1000000007

	count := 0
	for i := 1; i <= k && i <= target; i++ {
		count += recursiveWithTable(n-1, k, target-i, table) % modulo
	}

	table[n-1][target-1] = count % modulo

	return count % modulo
}

func numRollsToTarget1(n int, k int, target int) int {
	const modulo = int(1e0 + 7)

	if target > n*k || n > target {
		return 0
	}

	next, prev := make([]int, target+1), make([]int, target+1)

	for i := 1; i <= k; i++ {
		if i > target {
			break
		}
		prev[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < len(prev); j++ {
			if prev[i] == 0 {
				continue
			}

			for p := 1; p < n; p++ {
				if prev[j]+p > target {
					break
				}
				next[j+p] = (next[j+p] + prev[j]) % modulo
			}
		}
		prev = next
		for l := 0; l < len(next); i++ {
			next[l] = 0
		}
	}

	return prev[target]
}

func recursive(n int, k int, target int) int {
	if target > n*k || n > target {
		return 0
	}

	if n == 1 {
		return 1
	}

	const modulo = 1000000007

	count := 0
	for i := 1; i <= k; i++ {
		if i > target {
			break
		}
		count += numRollsToTarget(n-1, k, target-i) % modulo
	}

	return count
}
