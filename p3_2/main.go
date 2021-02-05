package p3_2

func Algo(nums []int, v int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] == v {
			result++
		}
	}
	return result
}
