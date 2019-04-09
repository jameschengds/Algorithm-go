package leetcode


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