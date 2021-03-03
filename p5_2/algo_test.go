package p5_2

import "testing"

func TestAlgo(t *testing.T) {
	a := []int{1, 3, 5, 7}

	r := Algo(a, 8)
	if !r {
		t.Fatal(r)
	}

	r = Algo(a, 12)
	if r {
		t.Fatal(r)
	}
}
