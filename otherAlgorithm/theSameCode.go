package otherAlgorithm

func TheSameNum(silceOne []int, sliceTwo []int) []int { //最简单的实现方式
	mapOne := make(map[int]bool, 0)
	sameNumSlice := make([]int, 0)
	for i := range silceOne {
		mapOne[silceOne[i]] = true
	}
	for i := range sliceTwo {
		_, ok := mapOne[sliceTwo[i]]
		if ok != false {
			sameNumSlice = append(sameNumSlice, sliceTwo[i])
		}
	}
	return sameNumSlice
}
func TheSameNumTwo(sliceOne []int, sliceTwo []int) []int { // 双指针的实现方式
	var i, j = 0, 0
	sameNumSlice := make([]int, 0)
	lenA := len(sliceOne)
	lenB := len(sliceTwo)
	for i < lenA && j < lenB {
		if sliceOne[i] < sliceTwo[j] {
			i++
		} else if sliceOne[i] == sliceTwo[j] {
			sameNumSlice = append(sameNumSlice, sliceOne[i])
			i++
			j++
		} else {
			j++
		}
	}
	return sameNumSlice
}
