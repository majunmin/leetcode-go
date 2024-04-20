package leetcode_2923

// https://leetcode.cn/problems/find-champion-i/
func findChampion(grid [][]int) int {
	// 排除 i == j 的 位置.
	// 某一行均为1 或者某一列均为0 则该队强
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	for i := 0; i < len(grid); i++ {
		if sum(grid[i]) == len(grid)-1 {
			return i
		}
	}
	return -1 // 不存在冠军
}

func sum(vals []int) int {
	var res int
	for _, num := range vals {
		res += num
	}
	return res
}
