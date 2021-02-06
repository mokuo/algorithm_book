package p3_4

import "math"

func Algo(nums []int) int {
	max := math.MinInt64
	min := math.MaxInt64

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	return max - min
}
