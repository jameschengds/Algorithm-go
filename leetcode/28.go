package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			ii := i
			for j := 0; j < len(needle); j++ {
				if haystack[ii] != needle[j] {
					break
				}
				ii++
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}
