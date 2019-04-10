package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	resp := make([][]int, 0)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					resp = append(resp, []int{nums[i], nums[j], nums[k]})
				} else if nums[i]+nums[j]+nums[k] > 0 {
					break
				}
			}
		}
	}
	//对接结果去重，这种可以成功但是leetcode会超时
	for i := range resp {
		for j := i + 1; j < len(resp); j++ {
			if resp[j][1] == resp[i][1] && resp[j][2] == resp[i][2] && resp[j][0] == resp[i][0] {
				resp = append(resp[:j], resp[j+1:]...)
				j--
			}
		}
	}
	return resp
}
