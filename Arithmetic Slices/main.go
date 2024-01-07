package main

func main() {

}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res := 0

	//minDiff := math.MaxInt
	//maxDiff := math.MinInt
	//for i := 0; i < len(nums)-1; i++ {
	//	if nums[i+1]-nums[i] > maxDiff {
	//		maxDiff = nums[i+1] - nums[i]
	//	}
	//	if nums[i+1]-nums[i] < minDiff {
	//		minDiff = nums[i+1] - nums[i]
	//	}
	//
	//}
	//
	//// We start checking for arithmetic subarrays
	//// We begin by checking for the existence of subarrays of length 3, then 4, and so on
	//currentLength := 3
	
	differences := make([]int, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		differences[i] = nums[i+1] - nums[i]
	}

	for i := 0; i < len(differences); i++ {
		currentDiff := differences[i]
		count := 0
		for i < len(differences) && differences[i] == currentDiff {
			count++
			i++
		}
		if count >= 2 {
			intermediate := 0
			for j := 1; j <= count-1; j++ {
				intermediate += j
			}
			res += intermediate
		}

		i--
	}

	return res
}
