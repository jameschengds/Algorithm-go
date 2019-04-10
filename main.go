package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	leet "github.com/Algorithm-go/leetcode"
)

func main() {
	//a := main1()
	//b := longestPalindrome("babad")
	//b := convert("Lawe", 2)
	//b := reverse(1234563)
	//b := niu("235KJ", "99TJQ")
	//fmt.Print(b)
	//b := niu("5JJJJ", "KKKK5")
	//b := leet.RomanToInt("LVIII")
	//var a = []string{"aca", "cba"}
	var a = []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}
	//b := leet.LongestCommonPrefix(a)
	b := leet.ThreeSum(a)
	fmt.Print(b)

}
func main1() []int {
	nums := [2]int{3, 3}
	target := 6
	mapp := make(map[int]int, 0)
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		mapp[nums[i]] = i
		var res = target - nums[i]
		if _, ok := mapp[res]; ok {
			if mapp[res] != i {
				if mapp[res] > i {
					ret = append(ret, i, mapp[res])
				} else {
					ret = append(ret, mapp[res], i)
				}
				return ret
			}
		}
	}
	return ret
}

//func mian2() int {
//	sByte := []byte("pwwkew")
//	var da,xiao int
//	var chang int
//	mapp := make(map[int]int, 0)
//	for i := 0; i < len(sByte); i++ {
//		if _, ok := mapp[int(sByte[i])]; ok {
//			da=
//			return i - mapp[int(sByte[i])]
//		}
//		mapp[int(sByte[i])] = i
//	}
//	return len(sByte)
//}
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

//对字符串进行分割，若n=4：每个unit的数量为6
//对应规律为2n-2，即s[0:6],s[6:13].等等
//我们可以对每一个unit的每一行设置一个切片，名字为s1,s2,s3,s4
//对第一个 i=0  unit处理的时候，s4=[s[n-1]],s3=[s[n-2],s[n]],s2=[s[n-3],s[n+1]],s1=[s[n-4]]
//对第二个 i=1  unit处理的时候，s4=[s[3],s[i*unit+3]],s3=[s[2],s[4]],s2=[s[1],s[5]],s1=[s[0]] n=4  j=0  l=3 r=3  j=1 l=2 r=4
func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	if numRows < 2 {
		return s
	}
	unit := 2*numRows - 2
	num := len(s)/unit + 1
	last := unit - len(s)%unit
	for i := 0; i < last; i++ {
		s = s + "#"
	}
	sByte := []byte(s)
	bucket := make([][]byte, numRows)
	for i := 0; i < num; i++ { //对每个unit进行操作
		for j := 0; j < numRows; j++ {
			str := make([]byte, 0)
			l := numRows - j - 1
			r := numRows + j - 1
			if l == r || (l-r)%unit == 0 {
				str = append(str, sByte[i*unit+numRows-j-1])
			} else {
				str = append(str, sByte[i*unit+numRows-j-1], sByte[i*unit+numRows+j-1])
			}
			bucket[j] = append(bucket[j], string(str)...)
		}
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		for _, v := range bucket[numRows-i-1] {
			result = append(result, v)
		}
	}
	//bb := []result
	aa := make([]byte, 0)
	for i, _ := range result {
		if string(result[i]) != "#" {
			aa = append(aa, result[i])
		}
	}
	return string(aa)
}

func reverse(x int) int {
	ans := 0
	for ; x != 0; x /= 10 {
		ans = ans*10 + x%10
	}

	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}

	return ans
}

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

func isPalindrome(x int) bool {
	var n int
	num := int64(x)
	if num < 0 {
		return false
	}
	for n = 0; num > 0; n++ {
		num /= 10
	}
	if n%2 == 1 { // ji shu
		//n/2+1 mid  52125
		for i := 0; i < n/2; i++ {
			r := num % int64(math.Pow10(i+1))
			l := num / int64(math.Pow10(n-i-1))
			if r != l {
				return false
			}
		}

	} else { //ou shu

	}
	return 1 == 1
}

func niu(paiA, paiB string) int {
	byteA := []byte(paiA)
	byteB := []byte(paiB)
	sliceA := make([]int, 0)
	sliceB := make([]int, 0)
	var flagA = false
	var flagB = false
	for i := range byteA {
		if byteA[i] > 57 {
			sliceA = append(sliceA, 10)
		} else {
			sliceA = append(sliceA, int(byteA[i]-48))
		}
	}
	for i := range byteB {
		if byteB[i] > 57 {
			sliceB = append(sliceB, 10)
		} else {
			sliceB = append(sliceB, int(byteB[i]-48))
		}
	}
	sort.Sort(sort.IntSlice(sliceA))
	sort.Sort(sort.IntSlice(sliceB))
	var sumA int
	var niuA int
	var sumB int
	var niuB int

	for _, v := range sliceA {
		if v == 10 {
			niuA += 1
		}
		sumA = sumA + v
	}
	for _, v := range sliceB {
		if v == 10 {
			niuB += 1
		}
		sumB = sumB + v
	}
	if niuA == 5 {
		flagA = true
	} else if niuA == 4 {
		flagA = true
	} else if niuA == 3 {
		flagA = true
	} else if niuA == 2 {
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j]+sliceA[k] == 10 {
						flagA = true
					}
				}
			}
		}
	} else if niuA == 1 { //2 or 3 ++ ==10
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				if v+sliceA[j] == 10 {
					flagA = true
				}
			}
		}
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j] == 10 {
						flagA = true
					}
				}
			}
		}
	} else if niuA == 0 {
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j] == 10 {
						flagA = true
					}
				}
			}
		}
	}
	if niuB == 5 {
		flagB = true
	} else if niuB == 4 {
		flagB = true
	} else if niuB == 3 {
		flagB = true
	} else if niuB == 2 {
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j]+sliceB[k] == 10 {
						flagB = true
					}
				}
			}
		}
	} else if niuB == 1 { //2 or 3 ++ ==10
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				if v+sliceB[j] == 10 {
					flagB = true
				}
			}
		}
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j] == 10 {
						flagB = true
					}
				}
			}
		}
	} else if niuB == 0 {
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j] == 10 {
						flagB = true
					}
				}
			}
		}
	}

	var leastA, leastB int
	if flagA == true && flagB == true {
		leastA = sumA % 10
		leastB = sumB % 10
		if leastA == 0 && leastB == 0 {
			return 0
		} else if leastA == 0 {
			return 1
		} else if leastB == 0 {
			return -1
		} else if leastA > leastB {
			return 1
		} else if leastA == leastB {
			return 0
		} else if leastA < leastB {
			return -1
		}
	} else if flagA == true {
		return 1
	} else if flagB == true {
		return -1
	} else if sliceA[0] == 1 && sliceB[0] != 1 {
		return 1
	} else if sliceA[0] != 1 && sliceB[0] == 1 {
		return -1
	} else if sliceA[0] == 1 && sliceB[0] == 1 {
		return 0
	} else if sliceA[4] > sliceB[4] {
		return 1
	} else if sliceA[4] < sliceB[4] {
		return -1
	} else if sliceA[4] == sliceB[4] {
		return 0
	}
	return 0
}
