package c5_7

import (
	"log"
)

func Algo(weight, value []int, W int) [][]int {
	N := len(weight)
	if N != len(value) {
		log.Fatal("weight と value の数が違う！！")
	}

	// 二重配列
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W)
	}

	for i := 0; i < N; i++ {
		for w := 0; w < W; w++ {
			// i 番目を選ぶケース
			if w-weight[i] >= 0 {
				chmax(&dp[i+1][w], dp[i][w-weight[i]]+value[i])
			}

			// i 番目を選ばないケース
			chmax(&dp[i+1][w], dp[i][w])
		}
	}
	return dp
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}
