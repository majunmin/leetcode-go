package leetcode_0051

import "strings"

// https://leetcode.cn/problems/n-queens/
// https://zhuanlan.zhihu.com/p/22846106
func solveNQueens(n int) [][]string {
	return solveNQueens1(n)

}

func solveNQueens1(n int) [][]string {
	if n < 0 {
		return nil
	}
	var result [][]string
	visitedRow := make(map[int]bool)
	visitedCol := make(map[int]bool)
	visitedPie := make(map[int]bool)
	visitedNa := make(map[int]bool)

	var bt func(curRow int, path []string)
	bt = func(curRow int, path []string) {
		if curRow == n {
			result = append(result, path)
			return
		}

		for i := 0; i < n; i++ {
			if visitedRow[curRow] || visitedCol[i] || visitedPie[i+curRow] || visitedNa[i-curRow+n] {
				continue
			}
			path = append(path, buildRow(i, n))
			visitedRow[curRow] = true
			visitedCol[i] = true
			visitedPie[i+curRow] = true
			visitedNa[i-curRow+n] = true
			bt(curRow+1, path)
			path = path[:len(path)-1]
			visitedRow[curRow] = false
			visitedCol[i] = false
			visitedPie[i+curRow] = false
			visitedNa[i-curRow+n] = false
		}
	}

	bt(0, []string{})
	return result
}

func backTraceSolution1(n int) [][]string {
	if n <= 0 {
		return nil
	}
	var result [][]string
	visitedRow := make(map[int]bool)
	visitedCol := make(map[int]bool)
	visitedPie := make(map[int]bool)
	visitedNa := make(map[int]bool)
	backTrace(&result, n, 0, []string{}, visitedRow, visitedCol, visitedPie, visitedNa)
	return result
}

func backTrace(result *[][]string, n int, curRow int, path []string, visitedRow map[int]bool, visitedCol map[int]bool, visitedPie map[int]bool, visitedNa map[int]bool) {
	//terminate
	if curRow == n {
		dst := make([]string, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	// for choice in choiceList
	for i := 0; i < n; i++ {
		if visitedRow[curRow] || visitedCol[i] || visitedPie[i+curRow] || visitedNa[i-curRow+n] {
			continue
		}
		path = append(path, buildRow(i, n))
		visitedRow[curRow] = true
		visitedCol[i] = true
		visitedPie[i+curRow] = true
		visitedNa[i-curRow+n] = true
		backTrace(result, n, curRow+1, path, visitedRow, visitedCol, visitedPie, visitedNa)
		visitedRow[curRow] = false
		visitedCol[i] = false
		visitedPie[i+curRow] = false
		visitedNa[i-curRow+n] = false
		path = path[:len(path)-1]
	}

}

func buildRow(i, n int) string {
	str := strings.Repeat(".", n)
	bytes := []byte(str)
	bytes[i] = 'Q'
	return string(bytes)
}

// 位运算优化 1
//
//	//把第 i 位置 1：a |= (1 << i)
//	//把第 i 位置 0：a &= ~(1 << i)
//	//把第 i 位取反：a ^= (1 << i)
//	//读取第 i 位的值：(a >> i) & 1
func solveNQueens2(n int) [][]string {
	if n < 0 {
		return nil
	}

	var result [][]string
	col, pie, na := 0, 0, 0
	backTraceSolution2(&result, n, col, pie, na, []string{}, 0)
	return result
}

func backTraceSolution2(result *[][]string, n int, colVisited int, pieVisited int, naVisited int, path []string, rowIdx int) {

	if rowIdx == n {
		dst := make([]string, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	// for choice In  choiceList
	for i := 0; i < n; i++ {
		//  剪枝
		pieIdx := rowIdx + i
		naIdx := i - rowIdx + n - 1
		if colVisited>>i&1 == 1 || pieVisited>>pieIdx&1 == 1 || naVisited>>naIdx&1 == 1 {
			continue
		}

		path = append(path, buildRow(n, i))
		colVisited |= 1 << i
		pieVisited |= 1 << pieIdx
		naVisited |= 1 << naIdx
		backTraceSolution2(result, n, colVisited, pieVisited, naVisited, path, rowIdx+1)
		path = path[:len(path)-1]
		colVisited &= ^(1 << i)
		pieVisited &= ^(1 << pieIdx)
		naVisited &= ^(1 << naIdx)
	}
}
