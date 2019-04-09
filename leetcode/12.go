package leetcode

func intToRoman(num int) string {
	var allNum =[]int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	var allChar= []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	var res string
	for i:=0;i<13;i++{
		for num>=allNum[i]{
			num=num-allNum[i]
			res=res+allChar[i]
		}
	}
	return res
}