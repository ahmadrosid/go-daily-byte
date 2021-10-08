package hammingdistance

import "fmt"

func Solve(x, y int) int {
	x1 := fmt.Sprintf("%08b", x)
	y1 := fmt.Sprintf("%08b", y)

	res := 0
	for i := 0; i < 8; i++ {
		if x1[i] != y1[i] {
			res++
		}
	}

	return res
}
