package leetcode

func maxArea(height []int) int {
	var max = 0
	var l =0
	var r = len(height)-1
	for l<r{
		var min int
		if height[l]>height[r]{
			min=height[r]
		}else {
			min=height[l]
		}
		s:=min*(r-l)
		if max>s{
			max=max
		}else {
			max=s
		}
		if height[l]<height[r]{
			l++
		}else{
			r--
		}
	}
	return max
}