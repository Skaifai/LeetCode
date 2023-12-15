package main

func main() {
	slices := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	print(destCity(slices))

}

func destCity(paths [][]string) string {
	destCities := make([]string, 0)
	for i := 0; i < len(paths); i++ {
		destCities = append(destCities, paths[i][1])
	}
	for i := 0; i < len(destCities); i++ {
		for j := 0; j < len(paths); j++ {
			if paths[j][0] == destCities[i] {
				destCities = append(destCities[:i], destCities[i+1:]...)
				i--
				break
			}
		}
	}
	return destCities[0]
}
