package main

func main() {

}

// Minimum Changes To Make Alternating Binary String
func minOperations(s string) int {
	first := 0
	second := 0
	for i := 0; i < len(s); i++ {
		if s[i] != uint8('0'+i%2) {
			first++
		} else {
			second++
		}
	}
	if first < second {
		return first
	}

	return second
}
