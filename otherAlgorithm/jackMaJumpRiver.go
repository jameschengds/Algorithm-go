package otherAlgorithm

func location(m, n, l int) (int, int) {
	y := l/m + 1
	x := l % n
	return x, y
}

func possibleWays(m, n, l int) [][][]int {
	x, y := location(m, n, l)
	possible := make([][][]int, 0)
	if x-2 > 0 && y+1 <= n {
		possible = append(possible, [][]int{
			{
				x - 2, y + 1,
			},
		})
	}
	if x+2 <= m && y+1 <= n {
		possible = append(possible, [][]int{
			{
				x + 2, y + 1,
			},
		})
	}
	if x-1 > 0 && y+2 <= n {
		possible = append(possible, [][]int{
			{
				x - 1, y + 2,
			},
		})
	}
	if x+1 <= m && y+2 <= n {
		possible = append(possible, [][]int{
			{
				x + 1, y + 2,
			},
		})
	}
	return possible
}

func getMinChoose(m, n int, inputXY [][][]int, input []int) (int, int) {
	var minStart = 0
	var next int
	for i := range inputXY {
		x := inputXY[i][0][0]
		y := inputXY[i][0][1]
		l := input[(y-1)*m+x-1]
		if i == 0 {
			minStart = l
		}
		if l <= minStart {
			minStart = l
		}
		next = (y-1)*m + x
	}
	return minStart, next
}

func JackMa(input []int, m, n int) int {
	var minRiverPath int
	for i := 0; i < m; i++ {
		eachMinRiverPath := input[i]
		nextPath := possibleWays(m, n, i+1)
		for len(nextPath) > 0 {
			//这里应该用递归算法，之前代码思路写的不对浪费了点时间。
			//下面的代码是错误的。
			minNextPath, next := getMinChoose(m, n, nextPath, input)
			eachMinRiverPath = eachMinRiverPath + minNextPath
			nextPath = possibleWays(m, n, next)
		}
		if i == 0 {
			minRiverPath = eachMinRiverPath
		}
		if minRiverPath > eachMinRiverPath {
			minRiverPath = eachMinRiverPath
		}
	}

	return minRiverPath
}
