package leetcode

func MaxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	res := nums[0]
	for i, v := range nums {
		if i > 0 {
			max1 := max * v
			min1 := min * v
			max = Max(Max(max1, min1), v)
			min = Min(Min(max1, min1), v)
			if res < max {
				res = max
			}
		}
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
