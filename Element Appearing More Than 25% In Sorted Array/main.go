package main

func main() {
	slice := []int{1, 2, 3, 3}
	print(findSpecialInteger(slice))
}

func findSpecialInteger(arr []int) int {
	threeQuarters := int(0.25 * float64(len(arr)))

	counter := 1

	currentNumber := arr[0]
	for i := 1; i < len(arr); i++ {
		for arr[i] == currentNumber {
			counter++
			i++
			if counter > threeQuarters {
				return currentNumber
			}
		}
		counter = 1
		currentNumber = arr[i]
	}
	return arr[len(arr)-1]
}
