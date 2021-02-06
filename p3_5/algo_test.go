package p3_5

import "testing"

func TestAlgo(t *testing.T) {
	var r int

	r = Algo([]int{1, 2})
	if r != 0 {
		t.Fatal(r)
	}

	r = Algo([]int{2, 4})
	if r != 1 {
		t.Fatal(r)
	}

	r = Algo([]int{1024, 2048, 4096})
	if r != 10 {
		t.Fatal(r)
	}
}
