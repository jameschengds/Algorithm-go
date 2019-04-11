package leetcode

import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	closestNum := nums[0] + nums[1] + nums[2]
	for i := 0; i < length-2; i++ {
		l := i + 1
		r := length - 1
		for l < r {
			three := nums[i] + nums[l] + nums[r]
			if calcAbs(three-target) < calcAbs(closestNum-target) {
				closestNum = three
			}
			if three > target {
				r--
			} else if three < target {
				l++
			} else {
				return target
			}
		}
	}
	return closestNum
}

func calcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}
