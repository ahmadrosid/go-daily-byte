package minimizepath

import "testing"

func Test_Solve(t *testing.T) {
	paths := []Path{
		{1, 1, 3},
		{2, 3, 1},
		{4, 6, 1},
	}
	tests := []struct {
		Paths    []Path
		Expected string
	}{
		{paths, "1->1->3->1->1"},
	}
	for _, data := range tests {
		res := Solve(data.Paths)
		if res != data.Expected {
			t.Fatalf("expected: '%s' got: '%s'", data.Expected, res)
		}
	}
}
