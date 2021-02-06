package p3_7

import (
	"log"
	"testing"
)

func TestAlgo(t *testing.T) {
	r := Algo("123")
	if r != 168 {
		log.Fatal(r)
	}

	r = Algo("1234")
	if r != 1736 {
		log.Fatal(r)
	}
}

func Test_insertPlus(t *testing.T) {
	r := insertPlus("01234", 3)
	if r != "012+34" {
		t.Fatal(r)
	}
}
