package leetcode

func LetterCombinations(digits string) []string {
	num := make(map[uint8][]string)
	num[2] = []string{"a", "b", "c"}
	num[3] = []string{"d", "e", "f"}
	num[4] = []string{"g", "h", "i"}
	num[5] = []string{"j", "k", "l"}
	num[6] = []string{"m", "n", "o"}
	num[7] = []string{"p", "q", "r", "s"}
	num[8] = []string{"t", "v", "u"}
	num[9] = []string{"w", "x", "y", "z"}
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var tmpStr string
	dfs(digits, &res, tmpStr, num, 0)
	return res
}

func dfs(digits string, res *[]string, tmpStr string, num map[uint8][]string, start int) {
	if start == len(digits) {
		*res = append(*res, tmpStr)
		return
	}
	a := digits[start] - 48
	tmp := num[a]
	for i := 0; i < len(tmp); i++ {
		dfs(digits, res, tmpStr+tmp[i], num, start+1)
	}
}
