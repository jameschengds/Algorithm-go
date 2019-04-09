package leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxLen := 0
	offset := 0
	for _, ch := range s {

		if seq := strings.IndexRune(s[left:right], ch); seq >= 0 {
			length := right - left
			if maxLen < length {
				maxLen = length
			}
			left = seq + offset + 1
			offset += seq + 1
		}
		right++
	}
	length := right - left
	if maxLen < length {
		maxLen = length
	}

	return maxLen
}