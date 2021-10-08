package traveltopoints

func Solve(points [][]int) int {
	lenX, lenY := len(points[0]), len(points)
	matrix := make([][]int, lenX)

	for x := 1; x < lenX; x++ {
		matrix[x] = make([]int, lenY)
		matrix[x][0] = points[x][0]
	}

	return 0
}
