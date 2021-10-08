package traveltopoints

import "testing"

func Test_Solve(t *testing.T) {
	point1 := [][]int{
		{0, 0}, {1, 1}, {2, 2},
	}
	point2 := [][]int{
		{0, 1}, {2, 3}, {4, 0},
	}
	tests := []struct {
		Paths    [][]int
		Expected int
	}{
		{point1, 2},
		{point2, 5},
	}
	for _, data := range tests {
		res := Solve(data.Paths)
		if res != data.Expected {
			t.Fatalf("expected: '%d' got: '%d'", data.Expected, res)
		}
	}
}
