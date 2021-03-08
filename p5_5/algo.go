package p5_5

// 配列 a（N個の正の整数）から、それぞれ何回でも足して良いとき、総和を W に一致させられるかを判定する
func Algo(a []int, W int) bool {
	N := len(a)

	// dp[i][j] => i番目までの a から、それぞれ何回でも足して良い時、 j に一致させられるか
	dp := make([][]bool, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]bool, W+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 {
				dp[i][j] = j%a[0] == 0
				continue
			}

			// a[i] を選ぶ時
			if j-a[i] >= 0 {
				chbool(&dp[i][j], dp[i-1][j-a[i]])
				chbool(&dp[i][j], dp[i][j-a[i]])
			}

			// a[i] を選ばない時
			chbool(&dp[i][j], dp[i-1][j])
		}
	}

	return dp[N-1][W]
}

func chbool(a *bool, b bool) {
	// a が false で b が true の時のみ true を代入する
	if !*a && b {
		*a = true
	}
}
