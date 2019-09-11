package otherAlgorithm

// 冒泡排序
func MaoPao(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var flag = false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return nums
}

//快速排序
func QuickSort(nums []int) []int {
	subQuickSort(nums, 0, len(nums)-1)
	return nums
}

func subQuickSort(nums []int, l, r int) {
	if l > r {
		return
	}
	q := partition(nums, l, r)
	subQuickSort(nums, l, q-1)
	subQuickSort(nums, q+1, r)
}

func partition(nums []int, l, r int) int {
	target := nums[r]
	t := r
	for l < r {
		for l < r && nums[l] <= target {
			l++
		}
		for l < r && nums[r] >= target {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[t] = nums[t], nums[l]
	return l
}
