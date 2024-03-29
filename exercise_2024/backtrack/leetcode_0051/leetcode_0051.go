package leetcode_0051

import "strings"

// https://leetcode.cn/problems/n-queens/
func solveNQueens(n int) [][]string {
	var result [][]string
	visitedCol := make([]bool, n)
	visitedPie := make([]bool, 2*n-1)
	visitedNa := make([]bool, 2*n-1)
	backtrace(&result, n, visitedCol, visitedPie, visitedNa, []string{}, 0)
	return result
}

func backtrace(result *[][]string, n int, visitedCol []bool, visitedPie []bool, visitedNa []bool, path []string, level int) {
	// terminate
	if level == n {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// backtrace
	// for choice in choiceList
	for i := 0; i < n; i++ {
		if visitedCol[i] || visitedPie[level+i] || visitedNa[level-i+n] {
			continue
		}
		visitedCol[i] = true
		visitedPie[level+i] = true
		visitedNa[level-i+n] = true

		path = append(path, buildPath(i, n))
		backtrace(result, n, visitedCol, visitedPie, visitedNa, path, level+1)
		path = path[:len(path)-1]

		visitedCol[i] = false
		visitedPie[level+i] = false
		visitedNa[level-i+n] = false
	}
}

func buildPath(idx int, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i == idx {
			sb.WriteByte('Q')
			continue
		}
		sb.WriteByte('.')
	}
	return sb.String()
}
