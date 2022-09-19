package offer_047

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/?envType=study-plan&id=lcof
func maxValue(grid [][]int) int {
	height := len(grid)
	if height == 0 {
		return -1
	}
	width := len(grid[0])
	if width == 0 {
		return -1
	}

	for i := 1; i < height; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < width; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			grid[i][j] += maxInt(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[height-1][width-1]

}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
