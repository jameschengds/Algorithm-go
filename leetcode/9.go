package leetcode

func isPalindrome(x int) bool {
	//	当x为负数，10的倍数时不可能是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	return x == reverses(x)
}

func reverses(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res = res*10 + x%10
	}
	return res
}