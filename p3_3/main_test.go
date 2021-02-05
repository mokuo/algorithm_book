package p3_3

import (
	"log"
	"testing"
)

func TestAlgo(t *testing.T) {
	r := Algo([]int{1, 2, 3, 4, 5})
	if r != 2 {
		log.Fatal("!!!")
	}

	r = Algo([]int{100, 2, 20, 80, 56})
	if r != 20 {
		log.Fatal("!!!")
	}
}
