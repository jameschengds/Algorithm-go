package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][3]int {
	sort.Ints(nums)
	length := len(nums)
	resp := make(map[[3]int]bool, 0)
	//response := make([][3]int, 0)
	re := make([][3]int, 0)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if _, ok := resp[[3]int{nums[i], nums[j], nums[k]}]; !ok {
						resp[[3]int{nums[i], nums[j], nums[k]}] = true
					}
				} else if nums[i]+nums[j]+nums[k] > 0 {
					break
				}
			}
		}
	}
	for i, _ := range resp {
		re = append(re, i)
	}
	return re
}
