package p5_4

import "testing"

func TestAlgo(t *testing.T) {
	a := []int{1, 3, 5, 7}

	r := Algo(a, 2, 8)
	if !r {
		t.Fatal(r)
	}

	r = Algo(a, 1, 8)
	if r {
		t.Fatal(r)
	}
}
