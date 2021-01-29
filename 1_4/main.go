package main

import (
	"errors"
	"fmt"
	"log"
)

//   a  6桁       xxxxxx   a6 a5 a4 a3 a2 a1
// x b  4桁     x   xxxx         b4 b3 b2 b1
// ---        ----------
//   c  6桁       66xxxx
//   d  6桁      6xxxxx
//   e  7桁    xx666xx
//   f  6桁    xx6xx6
// ---        ----------
//   g 10桁   xxxx66xxxx

func main() {
	for a := 100000; a < 1000000; a++ {
		for b := 1000; b < 10000; b++ {
			if (a%100000 == 0) && (b%1000 == 0) {
				fmt.Printf("a: %d, b: %d\n", a, b)
			}

			r, err := calc(a, b)
			if err != nil {
				fmt.Printf("a: %d, b: %d\n", a, b)
				log.Fatal(err)
			}

			if r == true {
				fmt.Printf("a: %d, b: %d\n", a, b)
				fmt.Println("sucess!")
				break
			}
		}
	}

	fmt.Println("該当なし。")
}

func calc(a, b int) (bool, error) {
	if !isNdigit(a, 6) {
		return false, errors.New("a が6桁じゃない")
	}
	if !isNdigit(b, 4) {
		return false, errors.New("b が4桁じゃない")
	}

	// c
	b1 := nDigit(b, 1)
	c := a * b1
	if !isNdigit(c, 6) {
		return false, nil
	}
	c5 := nDigit(c, 5)
	if c5 != 6 {
		return false, nil
	}
	c6 := nDigit(c, 6)
	if c6 != 6 {
		return false, nil
	}

	// d
	b2 := nDigit(b, 2)
	d := a * b2
	if !isNdigit(d, 6) {
		return false, nil
	}
	d6 := nDigit(d, 6)
	if d6 != 6 {
		return false, nil
	}

	// e
	b3 := nDigit(b, 3)
	e := a * b3
	if !isNdigit(e, 7) {
		return false, nil
	}
	e3 := nDigit(e, 3)
	if e3 != 6 {
		return false, nil
	}
	e4 := nDigit(e, 4)
	if e4 != 6 {
		return false, nil
	}
	e5 := nDigit(e, 5)
	if e5 != 6 {
		return false, nil
	}

	// f
	b4 := nDigit(b, 4)
	f := a * b4
	if !isNdigit(f, 6) {
		return false, nil
	}
	f1 := nDigit(f, 1)
	if f1 != 6 {
		return false, nil
	}
	f4 := nDigit(f, 4)
	if f4 != 6 {
		return false, nil
	}

	// g
	g := c * (d * 10) * (e * 100) * (f * 1000)
	if !isNdigit(g, 10) {
		return false, nil
	}
	g5 := nDigit(g, 5)
	if g5 != 6 {
		return false, nil
	}
	g6 := nDigit(g, 6)
	if g6 != 6 {
		return false, nil
	}

	return true, nil
}

// a が n 桁の数値なら true を返す
func isNdigit(a, n int) bool {
	switch n {
	case 1:
		return a >= 1 && a < 10
	case 2:
		return a >= 10 && a < 100
	case 3:
		return a >= 1 && a < 1000
	case 4:
		return a >= 1000 && a < 10000
	case 5:
		return a >= 10000 && a < 100000
	case 6:
		return a >= 100000 && a < 1000000
	case 7:
		return a >= 1000000 && a < 10000000
	case 8:
		return a >= 10000000 && a < 100000000
	case 9:
		return a >= 100000000 && a < 1000000000
	case 10:
		return a >= 1000000000 && a < 10000000000
	default:
		log.Fatalf("想定外の桁です: %d", n)
		return false
	}
}

// a の n 桁目を返す
func nDigit(a, n int) int {
	switch n {
	case 1:
		return a % 10
	case 2:
		return a % 100 / 10
	case 3:
		return a % 1000 / 100
	case 4:
		return a % 10000 / 1000
	case 5:
		return a % 100000 / 10000
	case 6:
		return a % 1000000 / 100000
	case 7:
		return a % 10000000 / 1000000
	case 8:
		return a % 100000000 / 10000000
	case 9:
		return a % 1000000000 / 100000000
	case 10:
		return a % 10000000000 / 1000000000
	default:
		log.Fatalf("想定外の桁です: %d", n)
		return 0
	}
}
