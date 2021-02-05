package main

import (
	"errors"
	"fmt"
)

//   ab
// x cd
// ----
//  efg
//  hi
// ----
//  jkl

//   2b
// x cd
// ----
//  e3g
//  hi
// ----
//  j4l

const a int = 2

// 正解のときだけ true を返す
func calc1_3(b, c, d int) (bool, error) {
	ab := a*10 + b

	efg := ab * d
	if efg < 100 || efg > 999 {
		return false, errors.New("efg が3桁じゃない")
	}

	f := efg % 100 / 10
	if f != 3 {
		return false, errors.New("f が3じゃない")
	}

	hi := ab * c
	if hi < 10 || hi > 99 {
		return false, errors.New("hi が2桁じゃない")
	}

	jkl := efg + (hi * 10)
	if jkl < 100 || jkl > 999 {
		return false, errors.New("jkl が3桁じゃない")
	}

	k := jkl % 100 / 10
	if k != 4 {
		return false, errors.New("k が4じゃない")
	}

	return true, nil
}

func main() {
	for b := 0; b < 10; b++ {
		for c := 0; c < 10; c++ {
			for d := 0; d < 10; d++ {
				r, _ := calc1_3(b, c, d)
				if r == true {
					fmt.Println("b:", b, "c:", c, "d:", d)
					break
				}
			}
		}
	}
}

// % go run 1_3.go
// b: 7 c: 3 d: 5
