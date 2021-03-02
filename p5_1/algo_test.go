package p5_1

import "testing"

func TestAlgo(t *testing.T) {
	// 幸福度 {海で泳ぐ, 虫取りする, 宿題をする}
	h := [][3]int{
		{3, 2, 1},
		{1, 2, 3},
		{3, 2, 1},
		{1, 3, 2},
	}

	r := Algo(h)
	if r != 12 {
		t.Fatal(r)
	}
}
