package minimizepath

func MinOfTwo(l, r int) int {
	if l < r {
		return l
	}

	return r
}

func Solve(grid [][]int) int {
	lenX, lenY := len(grid[0]), len(grid)
	matrix := make([][]int, lenY)
	for sum, y := 0, 0; y < lenY; y++ {
		matrix[y] = make([]int, lenX)
		sum += grid[y][0]
		matrix[y][0] = sum
	}

	for sum, x := 0, 0; x < lenX; x++ {
		sum += grid[0][x]
		matrix[0][x] = sum
	}

	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			matrix[y][x] = MinOfTwo(matrix[y][x-1], matrix[y-1][x]) + grid[y][x]
		}
	}

	return matrix[lenY-1][lenX-1]
}
