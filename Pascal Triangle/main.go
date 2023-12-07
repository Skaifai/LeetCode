package main

func main() {

}

func generate(numRows int) [][]int {
	triangle := make([][]int, 0)
	triangle = append(triangle, []int{1})
	for i := 1; i < numRows; i++ {
		line := make([]int, i+1)
		line[0] = 1
		line[i] = 1

		for j := 1; j < i; j++ {
			line[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle = append(triangle, line)
	}
	return triangle
}
