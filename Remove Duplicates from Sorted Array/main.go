package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	count := len(nums)
	for i := 0; i < count-1; i++ {
		if nums[i] == nums[i+1] {
			for j := i; j < count-1; j++ {
				nums[j] = nums[j+1]
			}
			i--
			count--
		}
	}
	return count
}
