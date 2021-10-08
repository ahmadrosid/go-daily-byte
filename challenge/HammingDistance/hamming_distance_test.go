package hammingdistance

import "testing"

func Test_Solve(t *testing.T) {
	x := 2
	y := 4
	expected := 4
	res := Solve(x, y)
	if expected != res {
		t.Fatalf("expected %d go %d", expected, res)
	}
}
