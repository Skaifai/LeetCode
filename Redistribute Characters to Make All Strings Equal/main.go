package main

func main() {

}

func makeEqual(words []string) bool {
	letters := make(map[rune]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			letters[rune(word[i])]++
		}
	}

	for _, letter := range letters {
		if letter%len(words) != 0 {
			return false
		}
	}

	return true
}
