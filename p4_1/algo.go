package p4_1

func Algo(N int) int {
	switch N {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	default:
		return Algo(N-1) + Algo(N-2) + Algo(N-3)
	}
}
