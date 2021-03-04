package p5_3

// a からいくつか選んで、総和を W に一致させられるかを判定する
func Algo(a []int, W int) int {
	N := len(a)

	// dp[i][j] => i番目までの a からいくつか選んで j に一致させられるかのブール値
	// a[i]を選ぶ時 => dp[i][j] は dp[i-1][j-a[i]] が true なら true
	// a[i]を選ばない時 => dp[i][j] は dp[i-1][j] が true なら true
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, W+1)
	}

	dp[0][0] = true
	dp[0][a[0]] = true

	for i := 1; i < N; i++ {
		for j := 0; j < W+1; j++ {
			if j-a[i] >= 0 {
				// a[i]を選ぶ時
				chbool(&dp[i][j], dp[i-1][j-a[i]])
			}
			// a[i]を選ばない時
			chbool(&dp[i][j], dp[i-1][j])
		}
	}

	cnt := 0
	for j := 1; j <= W; j++ {
		if dp[N-1][j] {
			cnt++
		}
	}

	return cnt
}

func chbool(a *bool, b bool) {
	// a が false で b が true の時のみ true を代入する
	if !*a && b {
		*a = true
	}
}
