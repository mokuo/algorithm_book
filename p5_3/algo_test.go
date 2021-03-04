package p5_3

import "testing"

func TestAlgo(t *testing.T) {
	a := []int{1, 3, 5, 7}

	r := Algo(a, 8)
	if r != 7 {
		t.Fatal(r)
	}

	r = Algo(a, 5)
	if r != 4 {
		t.Fatal(r)
	}
}
