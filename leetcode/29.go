package leetcode

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	var flag = 1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		flag = 1
	} else {
		flag = 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	dividend64 := int64(dividend)
	if dividend64 < 0 {
		dividend64 = -1 * dividend64
	}
	divisor64 := int64(divisor)
	if divisor64 < 0 {
		divisor64 = -1 * divisor64
	}

	var result = 0
	for i := 31; i >= 0; i-- {
		ii := uint(i)
		if (dividend64 >> ii) >= divisor64 { //找出足够大的数2^n*divisor
			result += 1 << ii             //将结果加上2^n
			dividend64 -= divisor64 << ii //将被除数减去2^n*divisor
		}
	}
	if flag == 0 {
		result = -1 * result
	}
	return result
}
