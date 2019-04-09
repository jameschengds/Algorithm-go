package leetcode


func longestPalindrome(s string) string {
	var str string
	for i := range s {
		len1 := len(compare(s, i, i))
		if len1 > len(str) {
			str = compare(s, i, i)
		}
		len2 := len(compare(s, i, i+1))
		if len2 > len(str) {
			str = compare(s, i, i+1)
		}
	}
	return str
}

func compare(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	str := s[l+1 : r]
	return str
}