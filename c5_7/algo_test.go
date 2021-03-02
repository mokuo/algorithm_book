package c5_7

import "testing"

func TestAlgo(t *testing.T) {
	wight := []int{2, 1, 3, 2, 1, 5}
	value := []int{3, 2, 6, 1, 3, 85}
	r := Algo(wight, value, 100)
	t.Log(r)
}
