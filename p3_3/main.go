package p3_3

// nums は2個以上で、全て異なる整数
// 2番目に小さい値を求める
func Algo(nums []int) int {
	var min1, min2 int

	if nums[0] < nums[1] {
		min1 = nums[0]
		min2 = nums[1]
	} else {
		min1 = nums[1]
		min2 = nums[2]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	return min2
}
