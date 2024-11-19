package leetcode_3240

import "fmt"

// https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/
func minFlips(grid [][]int) int {
	var (
		m, n   = len(grid), len(grid[0])
		result int
	)

	// 1. 统计 非中间行, 中间列的元素
	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			cnt1 := grid[i][j] + grid[i][n-j-1] + grid[m-i-1][j] + grid[m-i-1][n-j-1]
			result += min(cnt1, 4-cnt1)
		}
	}

	// 2. 统计中心位置的 元素
	// 中心位置必须是0, 否则 1 的个数是奇数.

	if m&1 == 1 && n&1 == 1 {
		result += grid[m/2][n/2]
	}

	var (
		diff, cnt1 int
	)
	// 3. 统计中间行
	if m&1 == 1 {
		for j := 0; j < n/2; j++ {
			if grid[m/2][j] != grid[m/2][n-1-j] {
				diff += 1
			} else {
				cnt1 += grid[m/2][j] * 2
			}
		}
	}
	// 4. 统计中间列
	if n&1 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2] != grid[m-i-1][n/2] {
				diff += 1
			} else {
				cnt1 += grid[i][n/2] * 2
			}
		}
	}
	fmt.Println(result)
	if diff > 0 {
		result += diff
	} else {
		result += cnt1 % 4
	}
	return result
}
