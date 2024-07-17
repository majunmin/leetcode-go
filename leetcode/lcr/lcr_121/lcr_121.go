package lcr_121

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findTargetIn2DPlants(plants [][]int, target int) bool {
	// param check
	if len(plants) == 0 || len(plants[0]) == 0 {
		return false
	}
	var (
		m, n     = len(plants), len(plants[0])
		row, col = 0, n - 1
	)
	for row < m && col >= 0 {
		if plants[row][col] == target {
			return true
		}
		if target < plants[row][col] {
			col--
		} else {
			row++
		}
	}
	return false
}
