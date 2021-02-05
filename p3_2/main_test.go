package p3_2

import "testing"

func TestAlgo(t *testing.T) {
	r := Algo([]int{1, 2, 3, 3, 4}, 3)
	if r != 2 {
		t.Fatal("!!!")
	}
}
