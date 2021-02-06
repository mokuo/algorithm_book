package p3_6

import "testing"

func TestAlgo(t *testing.T) {
	var r int

	r = Algo(4, 3)
	if r != 10 {
		t.Fatal(r)
	}

	r = Algo(3, 2)
	if r != 6 {
		t.Fatal(r)
	}

	r = Algo(2, 3)
	if r != 7 {
		t.Fatal(r)
	}
}
