package minimizepath

import "testing"

func Test_Solve(t *testing.T) {
	paths := [][]int{
		{1, 1, 3},
		{2, 3, 1},
		{4, 6, 1},
	}
	tests := []struct {
		Paths    [][]int
		Expected int
	}{
		{paths, 7},
	}
	for _, data := range tests {
		res := Solve(data.Paths)
		if res != data.Expected {
			t.Fatalf("expected: '%d' got: '%d'", data.Expected, res)
		}
	}
}
