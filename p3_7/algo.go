package p3_7

import (
	"go/constant"
	"go/token"
	"go/types"
	"log"
)

// ex) s => "1234"
func Algo(s string) int {
	sum := 0
	n := len(s) - 1 // + を入れられるのは n 箇所
	fset := token.NewFileSet()

	// n箇所に + を入れるか入れないか、を 1（入れる）と 0（入れない）に見立てる
	// n箇所1と0の組み合わせを10進数に変換して考える => bit
	// ビット論理積を使って 1 の箇所を取り出せる
	for bit := 0; bit < (1 << n); bit++ {
		expr := s

		// + を挿入する
		for i := n - 1; i >= 0; i-- {
			if bit&(1<<i) != 0 {
				expr = insertPlus(expr, i+1)
			}
		}

		// 式を評価する ex) "12+34" = 46
		tv, evalErr := types.Eval(fset, nil, token.NoPos, expr)
		if evalErr != nil {
			log.Fatal(evalErr)
		}

		r, _ := constant.Int64Val(tv.Value)
		sum = sum + int(r)
	}

	return sum
}

func insertPlus(s string, i int) string {
	before := s[:i]
	after := s[i:]

	return before + "+" + after
}
