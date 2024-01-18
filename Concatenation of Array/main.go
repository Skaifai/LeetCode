package main

func main() {

}

func getConcatenation(nums []int) []int {
	for i := 0; i < len(nums)-i; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}
