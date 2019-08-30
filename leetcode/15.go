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

func ThreeSum_1(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := [][]int{}
	var j int
	var k int
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j = i + 1
		k = length - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s < 0 {
				j = j + 1
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if s > 0 {
				k = k - 1
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}
