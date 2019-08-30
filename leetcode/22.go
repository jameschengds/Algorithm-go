package leetcode

func GenerateParenthesis(n int) []string {
	var res []string
	fn(&res, "", 0, 0, n)
	return res
}
func fn(res *[]string, str string, l, r, n int) {
	if l > n || r > n || r > l {
		return
	}
	if l == n && r == n {
		*res = append(*res, str)
	}
	fn(res, str+"(", l+1, r, n)
	fn(res, str+")", l, r+1, n)
}
