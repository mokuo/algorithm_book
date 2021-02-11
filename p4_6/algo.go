package p4_6

// Algo N 個の配列 A から何個か選んで、総和を W にできるか
func Algo(A []int, W int) bool {
	N := len(A)
	memo := make([][]bool, N+1)
	for i := range memo {
		memo[i] = make([]bool, W+1)
	}
	return algo(N, W, A, memo)
}

// a から i 個選んで、総和を w にできるか
func algo(i, w int, a []int, memo [][]bool) bool {
	// ベースケース
	if i == 0 {
		if w == 0 {
			return true
		}
		return false
	}

	if memo[i][w] {
		return memo[i][w]
	}

	// a[i - 1] を選ばない場合
	if algo(i-1, w, a, memo) {
		memo[i][w] = true
		return true
	}

	// a[i - 1] を選ぶ場合
	if algo(i-1, w-a[i-1], a, memo) {
		memo[i][w] = true
		return true
	}

	// どちらも false の場合は false
	memo[i][w] = false
	return false
}
