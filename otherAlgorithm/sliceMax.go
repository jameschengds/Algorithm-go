package otherAlgorithm

func FindMax(res []int) int {
	for i := 1; i < len(res); i++ {
		res[i] = max(res[i], res[i-1]+res[i])
	}
	max := res[0]
	for i := range res {
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
