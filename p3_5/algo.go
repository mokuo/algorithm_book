package p3_5

// nums はすべて正の整数
func Algo(nums []int) int {
	newNums := nums
	isAllPositive := true
	cnt := 0

	for {
		newNums, isAllPositive = algo(newNums)
		if isAllPositive == false {
			break
		}
		cnt++
	}

	return cnt
}

// すべて偶数なら 2 で割った値に置き換える
func algo(nums []int) ([]int, bool) {
	newNums := nums
	isAllPositive := true

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 { // 偶数なら2で割る
			newNums[i] = nums[i] / 2
		} else { // 奇数なら、元の配列に戻す
			isAllPositive = false
			newNums = nums
			break
		}
	}

	return newNums, isAllPositive
}
