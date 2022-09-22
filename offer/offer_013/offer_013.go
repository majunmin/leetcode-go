package offer_013

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

//https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/?envType=study-plan&id=lcof
func movingCount(m int, n int, k int) int {
	// param check
	if m == 0 || n == 0 {
		return 0
	}

	if k == 0 {
		return 1
	}
	//
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	return dsf(m, n, 0, 0, k, visited)
}

func dsf(m int, n int, rowIdx int, colIdx int, k int, visited [][]bool) int {

	if visited[rowIdx][colIdx] {
		return 0
	}
	if !check(rowIdx, colIdx, k) {
		return 0
	}

	count := 1
	visited[rowIdx][colIdx] = true
	for _, dir := range directions {
		newRow, newCol := rowIdx+dir[0], colIdx+dir[1]
		// check  row
		if 0 <= newRow && newRow < m && 0 <= newCol && newCol < n {
			count += dsf(m, n, newRow, newCol, k, visited)
		}
	}
	return count
}

func check(x, y, k int) bool {
	var sum int
	for x > 0 {
		sum += x % 10
		x = x / 10
	}
	for y > 0 {
		sum += y % 10
		y = y / 10
	}

	return sum <= k
}
