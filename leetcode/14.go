package leetcode

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	var common = strs[0]
	strByte := make([][]byte, 0)
	for i := range strs {
		bro := []byte(strs[i])
		strByte = append(strByte, bro)
	}
	for i := 1; i < length; i++ {
		var minlen int
		if len(strs[i]) > len(common) {
			minlen = len(common)
		} else {
			minlen = len(strs[i])
		}
		temp := make([]byte, 0)
		for j := 0; j < minlen; j++ {
			if common[j] == []byte(strs[i])[j] {
				temp = append(temp, common[j])
			} else {
				break
			}
		}
		common = string(temp)
	}
	return common
}
