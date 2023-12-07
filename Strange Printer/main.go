package main

import "fmt"

func main() {
	str := "abcdadcbaf"
	fmt.Println(strangePrinter(str))
}

func strangePrinter(s string) int {
	//r := 0
	type Interval struct {
		FirstAppearance int
		LastAppearance  int
	}
	type CharInfo struct {
		Char      rune
		Intervals []Interval
		Count     int
	}

	// Step 1: Initialize a map to track unique characters and their info
	charInfoMap := make(map[rune]*CharInfo)

	// Step 2: Iterate over the string and update character info in the map
	for idx, char := range s {
		if _, found := charInfoMap[char]; found {
			// Character already exists in the map, update last appearance index and increment count
			charInfoMap[char].Intervals[0].LastAppearance = idx
			charInfoMap[char].Count++
		} else {
			// Character is seen for the first time, create a new entry in the map
			charInfoMap[char] = &CharInfo{
				Char:  char,
				Count: 1,
			}
			charInfoMap[char].Intervals = make([]Interval, 0)
			charInfoMap[char].Intervals = append(charInfoMap[char].Intervals, Interval{
				FirstAppearance: idx,
				LastAppearance:  idx,
			})
		}
	}

	// Step 3: Convert the map to a slice to create the table
	charTable := make([]*CharInfo, 0, len(charInfoMap))
	for _, charInfo := range charInfoMap {
		charTable = append(charTable, charInfo)
	}

	// Step 4: Find all nested intervals for each letter. Intervals can be of length 1.
	for i := 0; i < len(charTable); i++ {
		leftBound := charTable[i].Intervals[0].FirstAppearance + 1
		rightBound := charTable[i].Intervals[0].LastAppearance - 1
		for leftBound < rightBound {
			intervalToAdd := Interval{
				FirstAppearance: 0,
				LastAppearance:  0,
			}

			for rune(s[leftBound]) != charTable[i].Char {
				leftBound++
			}
			if rune(s[leftBound]) == charTable[i].Char {
				intervalToAdd.FirstAppearance = leftBound
				leftBound++
				for rune(s[leftBound]) == charTable[i].Char {
					leftBound++
				}
				for rune(s[leftBound]) != charTable[i].Char && leftBound < rightBound {
					leftBound++
				}
			}

			for rune(s[rightBound]) != charTable[i].Char {
				rightBound--
			}
			if rune(s[rightBound]) == charTable[i].Char {
				intervalToAdd.LastAppearance = rightBound
				rightBound--
				for rune(s[rightBound]) == charTable[i].Char {
					rightBound--
				}
				for rune(s[rightBound]) != charTable[i].Char && rightBound > leftBound {
					rightBound--
				}
			}

			charTable[i].Intervals = append(charTable[i].Intervals, intervalToAdd)
		}

		if leftBound == rightBound && rune(s[leftBound]) == charTable[i].Char {
			intervalToAdd := Interval{
				FirstAppearance: leftBound,
				LastAppearance:  rightBound,
			}
			charTable[i].Intervals = append(charTable[i].Intervals, intervalToAdd)
		}
	}

	return 0
}
