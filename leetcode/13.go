package leetcode

func RomanToInt(s string) int {
	var num int
	var length = len(s)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "M":
			num = num + 1000
		case "C":
			if i != length-1 && string(s[i+1]) == "M" {
				num += 900
				i++
			} else if i != length-1 && string(s[i+1]) == "D" {
				num += 400
				i++
			} else {
				num += 100
			}
		case "D":
			num += 500
		case "X":
			if i != length-1 && string(s[i+1]) == "C" {
				num += 90
				i++
			} else if i != length-1 && string(s[i+1]) == "L" {
				num += 40
				i++
			} else {
				num += 10
			}
		case "I":
			if i != length-1 && string(s[i+1]) == "X" {
				num += 9
				i++
			} else if i != length-1 && string(s[i+1]) == "V" {
				num += 4
				i++
			} else {
				num += 1
			}
		case "L":
			num += 50
		case "V":
			num += 5

		}

	}

	return num

}
