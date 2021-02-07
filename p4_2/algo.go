package p4_1

type Val struct{ val int }

func Algo(N int) int {
	cache := map[int]int{}
	return algo(N, cache)
}

func algo(N int, cache map[int]int) int {
	switch N {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	default:
		val, ok := cache[N]
		if ok {
			return val
		}
		cache[N] = algo(N-1, cache) + algo(N-2, cache) + algo(N-3, cache)
		return cache[N]
	}
}
