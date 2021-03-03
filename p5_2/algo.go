package p5_2

// 0001
// 0010
// 0011
// 0100
// 0101
// 0110
// 0111
// ...

// a からいくつか選んで、総和を W に一致させられるかを判定する
func Algo(a []int, W int) bool {
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
			if j-a[i] < 0 {
				continue
			}
			// a[i]を選ぶ時
			chbool(&dp[i][j], dp[i-1][j-a[i]])
			// a[i]を選ばない時
			chbool(&dp[i][j], dp[i-1][j])
		}
	}

	return dp[N-1][W]
}

func chbool(a *bool, b bool) {
	// a が false の時のみ b を代入する
	if !*a {
		*a = b
	}
}
