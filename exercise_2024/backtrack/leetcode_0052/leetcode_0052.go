package leetcode_0052

func totalNQueens(n int) int {
	var result int
	backtrace(&result, n, 0, 0, 0, 0)
	return result
}

func backtrace(result *int, n int, visitedCol int, visitedPie int, visitedNa int, level int) {
	if n == level {
		*result = *result + 1
		return
	}
	// for choice in choiceList
	for i := 0; i < n; i++ {
		if visitedCol>>i == 1 || visitedNa>>(level+i) == 1 || visitedPie>>(level-i+n-1) == 1 {
			continue
		}
		visitedCol |= 1 << i
		visitedNa |= 1 << (level + i)
		visitedPie |= 1 << (level - i + n - 1)
		backtrace(result, n, visitedCol, visitedPie, visitedNa, level+1)
		visitedCol &= ^(1 << i)
		visitedNa &= ^(1 << (level + i))
		visitedPie &= ^(1 << (level - i + n - 1))
	}
}
