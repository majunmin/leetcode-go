package _024_02_04

var (
	directions = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{0, 0},
	}
)

// https://leetcode.cn/contest/weekly-contest-383/problems/find-the-grid-of-region-average/
func resultGrid(image [][]int, threshold int) [][]int {
	// param check
	if len(image) < 3 || len(image[0]) < 3 {
		// param invalid param
		panic("invalid param, image width or height smaller than 3")
	}
	m, n := len(image), len(image[0])

	cnt := make([][]int, m)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}
	// cal sum, build result sum
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !isRegion(image, i, j, threshold) {
				continue
			}
			var sum int
			for _, dir := range directions {
				sum += image[i+dir[0]][j+dir[1]]
			}
			avg := sum / 9
			for _, dir := range directions {
				result[i+dir[0]][j+dir[1]] += avg
				cnt[i+dir[0]][j+dir[1]] += 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cnt[i][j] > 0 {
				result[i][j] = result[i][j] / cnt[i][j]
			} else {
				result[i][j] = image[i][j]
			}
		}
	}

	return result
}

func isRegion(image [][]int, i int, j int, threshold int) bool {
	if abs(image[i][j]-image[i-1][j]) > threshold {
		return false
	}
	if abs(image[i][j]-image[i+1][j]) > threshold {
		return false
	}
	if abs(image[i][j]-image[i][j-1]) > threshold {
		return false
	}
	if abs(image[i][j]-image[i][j+1]) > threshold {
		return false
	}
	if abs(image[i-1][j]-image[i-1][j-1]) > threshold {
		return false
	}
	if abs(image[i-1][j]-image[i-1][j+1]) > threshold {
		return false
	}
	if abs(image[i+1][j]-image[i+1][j-1]) > threshold {
		return false
	}
	if abs(image[i+1][j]-image[i+1][j+1]) > threshold {
		return false
	}

	if abs(image[i][j-1]-image[i-1][j-1]) > threshold {
		return false
	}
	if abs(image[i][j-1]-image[i+1][j-11]) > threshold {
		return false
	}
	if abs(image[i][j+1]-image[i-1][j+1]) > threshold {
		return false
	}
	if abs(image[i][j+1]-image[i+1][j+1]) > threshold {
		return false
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
