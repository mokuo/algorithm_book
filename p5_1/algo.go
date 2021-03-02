package p5_1

import "math"

// 動的計画法
// 幸福度 {0:海で泳ぐ, 1:虫取りする, 2:宿題をする}
func Algo(h [][3]int) int {
	N := len(h)
	dp := make([][3]int, N)

	// dpテーブル初期化
	for i := 0; i < N; i++ {
		dp[i][0] = math.MinInt64
		dp[i][1] = math.MinInt64
		dp[i][2] = math.MinInt64
	}

	dp[0][0] = h[0][0]
	dp[0][1] = h[0][1]
	dp[0][2] = h[0][2]

	for i := 1; i < N; i++ {
		chmax(&dp[i][0], dp[i-1][1]+h[i][0])
		chmax(&dp[i][0], dp[i-1][2]+h[i][0])

		chmax(&dp[i][1], dp[i-1][0]+h[i][1])
		chmax(&dp[i][1], dp[i-1][2]+h[i][1])

		chmax(&dp[i][2], dp[i-1][0]+h[i][2])
		chmax(&dp[i][2], dp[i-1][1]+h[i][2])
	}

	return max(dp[N-1])
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func max(a [3]int) int {
	m := math.MinInt64
	for i := 0; i < 3; i++ {
		if m < a[i] {
			m = a[i]
		}
	}
	return m
}
