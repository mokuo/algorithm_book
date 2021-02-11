package p4_6

import "testing"

func TestAlgo(t *testing.T) {
	r := Algo([]int{3, 2, 6, 5}, 14)

	if !r {
		t.Fatal("!!!")
	}
}
