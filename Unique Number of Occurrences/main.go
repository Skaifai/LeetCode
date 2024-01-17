package main

func main() {

}

func uniqueOccurrences(arr []int) bool {
	occurrences := map[int]int{}
	for i := 0; i < len(arr); i++ {
		occurrences[arr[i]]++
	}
	unique := map[int]int{}
	for _, i := range occurrences {
		if unique[i] == 1 {
			return false
		}
		unique[i]++
	}
	return true
}
