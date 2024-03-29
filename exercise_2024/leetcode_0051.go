package exercise_2024

import "strings"

// https://leetcode.cn/problems/n-queens/?envType=study-plan-v2&envId=top-100-liked
func solveNQueens(n int) [][]string {

	// param check
	if n <= 0 {
		return nil
	}

	var result [][]string
	var col, pie, na int // 记录已经 防止过 皇后的位置
	backtrace(&result, n, 0, col, pie, na, []string{})
	return result
}

// pie : row + col
// na : col - row + n
// col : col
func backtrace(result *[][]string, n int, level int, col int, pie int, na int, path []string) {
	// terminate
	if level == n {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// for choice in choiceList
	for i := 0; i < n; i++ {
		if col>>i&1 == 1 || pie>>(i+level)&1 == 1 || na>>(i-level+n)&1 == 1 {
			continue
		}
		path = append(path, makeRow(n, i))
		col |= 1 << i
		pie |= 1 << (i + level)
		na |= 1 << (i - level + n)
		backtrace(result, n, level+1, col, pie, na, path)
		col &= ^(1 << i)
		pie &= ^(1 << (i + level))
		na &= ^(1 << (i - level + n))
		path = path[:len(path)-1]
	}
}

func makeRow(n, idx int) string {
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
