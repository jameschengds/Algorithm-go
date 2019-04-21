package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	rtIndex := -1
	end := len(nums) - 1
	begin := 0
	if nums[end] < target && nums[begin] > target {
		return -1
	}
	for begin < end {
		if nums[begin] == target {
			rtIndex = begin
			break
		}
		if nums[end] == target {
			rtIndex = end
			break
		}
		if nums[begin] < target && begin+1 < end && nums[begin+1] > nums[begin] {
			begin++
		} else if nums[end] > target && end-1 > begin && nums[end-1] < nums[end] {
			end--
		} else {
			break
		}
	}
	return rtIndex

}
