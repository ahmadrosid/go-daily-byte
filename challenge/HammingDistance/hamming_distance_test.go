package hammingdistance

import (
	"testing"
)

func Test_Solve(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{2, 4, 2},
		{1, 4, 2},
		{5, 4, 1},
	}

	for _, el := range tests {
		res := Solve(el.x, el.y)
		if el.expected != res {
			t.Logf("x: %d, y: %d, res: %d", el.x, el.y, res)
			t.Fatalf("expected %d got %d", el.expected, res)
		}
	}
}
