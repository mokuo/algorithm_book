package p4_5

import (
	"fmt"
	"strings"
)

// Algo K 以下の753数がいくつあるか。7,5,3 は最低1回は使う
func Algo(K int) int {
	results := map[int]bool{}

	append357(0, 3, K, results)
	append357(0, 5, K, results)
	append357(0, 7, K, results)

	// fmt.Println(results)
	return len(results)
}

func append357(num, n, K int, results map[int]bool) {
	// n は初回以外 3 or 5 or 7
	// num * 10 + n する
	result := num*10 + n

	// 結果が K より大きい => 何もせず終了させる
	if result > K {
		return
	}

	// 3,5,7をすべて含む場合、取っておく
	if isOK(result) {
		results[result] = true
	}

	// 再帰関数に 3, 5, 7 を渡す
	append357(result, 3, K, results)
	append357(result, 5, K, results)
	append357(result, 7, K, results)
}

// num に 3,5,7 がすべて含まれるか
func isOK(num int) bool {
	strNum := fmt.Sprint(num)
	return strings.Contains(strNum, "3") && strings.Contains(strNum, "5") && strings.Contains(strNum, "7")
}
