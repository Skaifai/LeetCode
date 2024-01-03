package main

func main() {

}

func numberOfBeams(bank []string) int {
	if len(bank) == 1 {
		return 0
	}
	result := 0
	prev := 0
	for _, row := range bank {
		count := 0
		for i := 0; i < len(row); i++ {
			if row[i] == '1' {
				count++
			}
		}
		if count != 0 {
			result += count * prev
			prev = count
		}
	}
	return result
}
