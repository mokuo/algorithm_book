package p4_1

import "testing"

func TestAlgo(t *testing.T) {
	r := Algo(3)
	if r != 1 {
		t.Fatal(r)
	}

	r = Algo(4)
	if r != 2 {
		t.Fatal(r)
	}

	r = Algo(5)
	if r != 4 {
		t.Fatal(r)
	}

	r = Algo(6)
	if r != 7 {
		t.Fatal(r)
	}

	r = Algo(7)
	if r != 13 {
		t.Fatal(r)
	}
}
