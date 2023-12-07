package main

func main() {
	s := "52"
	lastDigitString := largestOddNumber(s)
	print(lastDigitString)
}

//func largestOddNumber(num string) string {
//	currentLength := len(num)
//
//	for i := 0; i < len(num); i++ {
//		for j := 0; j < len(num)-currentLength+1; j++ {
//			substring := num[j : j+currentLength]
//			lastDigit := substring[len(substring)-1] - 48
//			if lastDigit%2 == 1 {
//				return substring
//			}
//		}
//		currentLength--
//	}
//	return ""
//}

func largestOddNumber(num string) string {
	for i := len(num); i > 0; i-- {
		lastDigit := num[i-1] - 48
		if lastDigit%2 == 1 {
			return num[0:i]
		}
	}

	return ""
}
