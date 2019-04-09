package leetcode

import (
	"math"
	"strconv"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var flag = 1
	var i, num int
	for i < len(str) && str[i] == 32 {
		i++
	}
	if i >= len(str) {
		return 0
	}
	if str[i] == 43 {
		i++
		//	-
	} else if str[i] == 45 {
		flag = -1
		i++
	}
	for ; i < len(str); i++ {
		//	后面为字符时，直接返回前面的数字
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) {
			return num * flag
		}
		//	获得每次计算的结果
		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		//	检查越界【越界就返回极限值】
		if num*flag < math.MinInt32 {
			return math.MinInt32
		} else if num*flag > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return num * flag
	//sByte := []byte(str)
	//for i := range str {
	//	if sByte[i] != 20 {
	//
	//	}
	//}
}