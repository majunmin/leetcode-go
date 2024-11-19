package leetcode_0661

var (
	directions = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
)

// https://leetcode.cn/problems/image-smoother/
func imageSmoother(img [][]int) [][]int {
	var (
		m, n   = len(img), len(img[0])
		result = make([][]int, m)
	)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	// cal
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := img[i][j]
			cnt := 1
			for _, dir := range directions {
				newI, newJ := i+dir[0], j+dir[1]
				if newI < 0 || newI >= m || newJ < 0 || newJ >= n {
					continue
				}
				sum += img[newI][newJ]
				cnt++
			}
			result[i][j] = sum / cnt
		}
	}
	return result
}
