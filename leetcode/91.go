package leetcode

import "strconv"

func NumDecodings(s string) int {
	dp := make([]int, 0)

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp = append(dp, 1)
	for i := 1; i < len(s); i++ {
		a, _ := strconv.Atoi(string(s[i-1]))
		b, _ := strconv.Atoi(string(s[i]))
		if a*10+b > 26 {
			dp = append(dp, dp[i-1])
		} else {
			dp = append(dp, dp[i-1]*2)
		}
	}
	return dp[len(s)-1]
}
