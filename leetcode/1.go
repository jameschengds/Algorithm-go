package leetcode

func leetcode1() []int {
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