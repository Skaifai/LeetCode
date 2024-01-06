package main

import (
	"fmt"
	"sort"
)

func main() {
	//startTime := []int{1, 1, 1}
	//endTime := []int{2, 3, 4}
	//profit := []int{5, 6, 4}
	//
	//startTime := []int{1, 2, 3, 3}
	//endTime := []int{3, 4, 5, 6}
	//profit := []int{50, 10, 40, 70}

	//startTime := []int{1, 2, 2, 3}
	//endTime := []int{2, 5, 3, 4}
	//profit := []int{3, 4, 1, 2}
	//
	//startTime := []int{4, 2, 4, 8, 2}
	//endTime := []int{5, 5, 5, 10, 8}
	//profit := []int{1, 2, 8, 10, 4}

	startTime := []int{6, 15, 7, 11, 1, 3, 16, 2}
	endTime := []int{19, 18, 19, 16, 10, 8, 19, 8}
	profit := []int{2, 9, 1, 19, 5, 7, 3, 19}
	fmt.Println(jobScheduling(startTime, endTime, profit))
}

type Job struct {
	StartTime int
	EndTime   int
	Profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]Job, 0)
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, Job{
			StartTime: startTime[i],
			EndTime:   endTime[i],
			Profit:    profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].StartTime < jobs[j].StartTime
	})
	fmt.Println(jobs)
	sort.Ints(startTime)
	dp := make([]int, len(startTime))
	result := recursion(startTime[0], startTime, jobs, dp)
	fmt.Println(dp)
	return result
}

func recursion(time int, startTime []int, jobs []Job, dp []int) int {
	index := sort.SearchInts(startTime, time)
	if index >= len(jobs) {
		return 0
	}
	if dp[index] != 0 {
		return dp[index]
	}
	intermediate := jobs[index].Profit
	for j := index; j < len(jobs) && startTime[j] >= startTime[index] && startTime[j] < jobs[index].EndTime; j++ {
		recRes := recursion(jobs[j].EndTime, startTime, jobs, dp) + jobs[j].Profit
		if recRes > intermediate {
			intermediate = recRes
		}
	}
	dp[index] = intermediate
	return dp[index]
}

//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func jobScheduling(startTime []int, endTime []int, profit []int) int {
//	priority := make([]float64, len(startTime))
//	start := math.MaxInt
//	end := math.MinInt
//	for i := 0; i < len(startTime); i++ {
//		priority[i] = float64(profit[i]) / float64(endTime[i]-startTime[i])
//		if startTime[i] < start {
//			start = startTime[i]
//		}
//		if endTime[i] > end {
//			end = endTime[i]
//		}
//	}
//	//sort.Slice(startTime, func(i, j int) bool {
//	//	return priority[i] > priority[j]
//	//})
//	//sort.Slice(endTime, func(i, j int) bool {
//	//	return priority[i] > priority[j]
//	//})
//	//sort.Slice(profit, func(i, j int) bool {
//	//	return priority[i] > priority[j]
//	//})
//	//sort.Slice(priority, func(i, j int) bool {
//	//	return priority[i] > priority[j]
//	//})
//	sort.Slice(endTime, func(i, j int) bool {
//		return startTime[i] < startTime[j]
//	})
//	sort.Slice(profit, func(i, j int) bool {
//		return startTime[i] < startTime[j]
//	})
//	sort.Slice(priority, func(i, j int) bool {
//		return startTime[i] < startTime[j]
//	})
//	sort.Slice(startTime, func(i, j int) bool {
//		return startTime[i] < startTime[j]
//	})
//	fmt.Println(priority)
//	fmt.Println(startTime)
//	fmt.Println(endTime)
//	fmt.Println(profit)
//
//	dp := make([]int, end-start+1)
//
//	result := recursion(start, startTime, endTime, profit, dp)
//
//	//for _, profitI := range dp {
//	//	result += profitI
//	//}
//
//	//for i := start; i <= end; {
//	//	index := sort.SearchInts(startTime, i)
//	//	if index >= len(profit) {
//	//		break
//	//	}
//	//	result += profit[index]
//	//	i = endTime[index]
//	//}
//
//	return result
//}
//
//func recursion(time int, startTime []int, endTime []int, profit []int, dp []int) int {
//
//	if time >= len(dp) {
//		return 0
//	}
//
//	//if dp[time] != 0 {
//	//	return dp[time]
//	//}
//	index := sort.SearchInts(startTime, time)
//	if index >= len(profit) {
//		return 0
//	}
//	intermediate := 0
//	//chosenIndex := index
//	recRes := 0
//	for j := index; j < len(profit) && startTime[j] >= startTime[index] && startTime[j] < endTime[index]; j++ {
//		recRes = recursion(endTime[j], startTime, endTime, profit, dp) + profit[j]
//		if recRes > intermediate {
//			intermediate = recRes
//			//chosenIndex = j
//		}
//	}
//	//for k := startTime[chosenIndex] - 1; k < endTime[chosenIndex]-1; k++ {
//	//	if intermediate > dp[k] {
//	//		dp[k] = intermediate
//	//	}
//	//}
//
//	//dp[time] = profit[chosenIndex]
//	return intermediate
//}
