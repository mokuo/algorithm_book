package p5_5

import "testing"

func TestAlgo(t *testing.T) {
	a := []int{3, 8, 2, 5}

	r := Algo(a, 18)
	if !r {
		t.Fatal(r)
	}

	r = Algo(a, 13)
	if !r {
		t.Fatal(r)
	}

	r = Algo(a, 1)
	if r {
		t.Fatal(r)
	}

	r = Algo(a, 17)
	if !r {
		t.Fatal(r)
	}

	a = []int{3, 8}

	r = Algo(a, 14)
	if !r {
		t.Fatal(r)
	}

	r = Algo(a, 19)
	if !r {
		t.Fatal(r)
	}
}
