package otherAlgorithm

import "sort"

func niu(paiA, paiB string) int {
	byteA := []byte(paiA)
	byteB := []byte(paiB)
	sliceA := make([]int, 0)
	sliceB := make([]int, 0)
	var flagA = false
	var flagB = false
	for i := range byteA {
		if byteA[i] > 57 {
			sliceA = append(sliceA, 10)
		} else {
			sliceA = append(sliceA, int(byteA[i]-48))
		}
	}
	for i := range byteB {
		if byteB[i] > 57 {
			sliceB = append(sliceB, 10)
		} else {
			sliceB = append(sliceB, int(byteB[i]-48))
		}
	}
	sort.Sort(sort.IntSlice(sliceA))
	sort.Sort(sort.IntSlice(sliceB))
	var sumA int
	var niuA int
	var sumB int
	var niuB int

	for _, v := range sliceA {
		if v == 10 {
			niuA += 1
		}
		sumA = sumA + v
	}
	for _, v := range sliceB {
		if v == 10 {
			niuB += 1
		}
		sumB = sumB + v
	}
	if niuA == 5 {
		flagA = true
	} else if niuA == 4 {
		flagA = true
	} else if niuA == 3 {
		flagA = true
	} else if niuA == 2 {
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j]+sliceA[k] == 10 {
						flagA = true
					}
				}
			}
		}
	} else if niuA == 1 { //2 or 3 ++ ==10
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				if v+sliceA[j] == 10 {
					flagA = true
				}
			}
		}
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j] == 10 {
						flagA = true
					}
				}
			}
		}
	} else if niuA == 0 {
		for i, v := range sliceA {
			for j := i + 1; j < len(sliceA); j++ {
				for k := j + 1; k < len(sliceA); k++ {
					if v+sliceA[j] == 10 {
						flagA = true
					}
				}
			}
		}
	}
	if niuB == 5 {
		flagB = true
	} else if niuB == 4 {
		flagB = true
	} else if niuB == 3 {
		flagB = true
	} else if niuB == 2 {
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j]+sliceB[k] == 10 {
						flagB = true
					}
				}
			}
		}
	} else if niuB == 1 { //2 or 3 ++ ==10
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				if v+sliceB[j] == 10 {
					flagB = true
				}
			}
		}
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j] == 10 {
						flagB = true
					}
				}
			}
		}
	} else if niuB == 0 {
		for i, v := range sliceB {
			for j := i + 1; j < len(sliceB); j++ {
				for k := j + 1; k < len(sliceB); k++ {
					if v+sliceB[j] == 10 {
						flagB = true
					}
				}
			}
		}
	}

	var leastA, leastB int
	if flagA == true && flagB == true {
		leastA = sumA % 10
		leastB = sumB % 10
		if leastA == 0 && leastB == 0 {
			return 0
		} else if leastA == 0 {
			return 1
		} else if leastB == 0 {
			return -1
		} else if leastA > leastB {
			return 1
		} else if leastA == leastB {
			return 0
		} else if leastA < leastB {
			return -1
		}
	} else if flagA == true {
		return 1
	} else if flagB == true {
		return -1
	} else if sliceA[0] == 1 && sliceB[0] != 1 {
		return 1
	} else if sliceA[0] != 1 && sliceB[0] == 1 {
		return -1
	} else if sliceA[0] == 1 && sliceB[0] == 1 {
		return 0
	} else if sliceA[4] > sliceB[4] {
		return 1
	} else if sliceA[4] < sliceB[4] {
		return -1
	} else if sliceA[4] == sliceB[4] {
		return 0
	}
	return 0
}
