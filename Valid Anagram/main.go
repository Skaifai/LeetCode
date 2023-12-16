package main

func main() {
	print(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[int(s[i])-'a']++
		freq[int(t[i])-'a']--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	hash := map[rune]int{}
//
//	for _, c := range s {
//		hash[c]++
//	}
//
//	for _, c := range t {
//		if hash[c] == 0 {
//			return false
//		}
//
//		hash[c]--
//	}
//
//	return true
//}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	for i := 0; i < len(s); i++ {
//		index := strings.Index(t, string(s[i]))
//		if index == -1 {
//			return false
//		} else {
//			t = t[:index] + t[index+1:]
//		}
//	}
//
//	if len(t) > 0 {
//		return false
//	}
//
//	return true
//}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	for i := 0; i < len(s); i++ {
//		for j := 0; j < len(t); j++ {
//			if s[i] == t[j] {
//				t = t[:j] + t[j+1:]
//				j--
//				break
//			}
//		}
//	}
//
//	if len(t) > 0 {
//		return false
//	}
//
//	return true
//}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	intS := make([]int, len(s))
//	intT := make([]int, len(t))
//
//	for i := 0; i < len(s); i++ {
//		intS[i] = int(s[i])
//		intT[i] = int(t[i])
//	}
//
//	sort.Ints(intS)
//	sort.Ints(intT)
//
//	for i := 0; i < len(s); i++ {
//		if intS[i] != intT[i] {
//			return false
//		}
//	}
//
//	return true
//}
