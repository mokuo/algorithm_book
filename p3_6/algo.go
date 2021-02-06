package p3_6

func Algo(k, n int) int {
	count := 0
	var z int

	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z = n - x - y
			if z >= 0 && z <= k {
				count++
			}
		}
	}

	return count
}
