package p3_4

import (
	"log"
	"testing"
)

func TestAlgo(t *testing.T) {
	var r int
	r = Algo([]int{-100, 0, 50, 100})
	if r != 200 {
		log.Fatal("!!!")
	}

	r = Algo([]int{20, 50, 1000, 1020})
	if r != 1000 {
		log.Fatal("!!!")
	}
}
