package p3_3

import (
	"math"
)

// nums は2個以上で、全て異なる整数
// 2番目に小さい値を求める
func Algo(nums []int) int {
	min1 := math.MaxInt64
	min2 := math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	return min2
}
