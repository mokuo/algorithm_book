package p5_4

// 配列 a から k 個以下の整数を選んで、総和を W に一致させられるかを判定する
func Algo(a []int, k, W int) bool {
	N := len(a)

	// dp[i][j] => i番目までの a からいくつ選んで j に一致させられるか
	// dp[i][j] = dp[i-1][j]
	// j == a[i] => dp[i][j] = 1
	// dp[i-1][j - a[j]] >= 0 => dp[i][j] = dp[i-1][j-a[i]] + 1
	// a[i]を選ぶ時 => dp[i][j] dp[i-1][j-a[i]] + 1
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, W+1)

		for j := 0; j <= W; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	dp[0][a[0]] = 1

	for i := 1; i < N; i++ {
		for j := 0; j < W+1; j++ {
			if j-a[i] >= 0 {
				// そのまま引き継ぐ
				dp[i][j] = dp[i-1][j]

				if dp[i-1][j-a[i]] >= 0 {
					// 0以上の小さい方を取る、とかした方が良さそう
					dp[i][j] = dp[i-1][j-a[i]] + 1
				}
			}
		}
	}

	return 0 <= dp[N-1][W] && dp[N-1][W] <= k
}
