package p4_5

import "testing"

func TestAlgo(t *testing.T) {
	r := Algo(800)
	if r != 6 {
		t.Fatal(r)
	}

	r = Algo(8000)
	if r != 42 {
		t.Fatal(r)
	}
}

func Test_isOK(t *testing.T) {
	r := isOK(3357)
	if !r {
		t.Fatal(r)
	}

	r = isOK(355)
	if r {
		t.Fatal(r)
	}

	r = isOK(77777)
	if r {
		t.Fatal(r)
	}
}
