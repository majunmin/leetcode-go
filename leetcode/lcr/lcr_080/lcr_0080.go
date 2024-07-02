package lcr_080

// https://leetcode.cn/problems/uUsW3B/
func combine(n int, k int) [][]int {
	var (
		result = make([][]int, 0)
	)
	backtace(&result, []int{}, k, n, 1)
	return result
}

func backtace(result *[][]int, path []int, k int, n int, begin int) {
	// termiante
	if k == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// for choice in choiceList
	for i := begin; i <= n; i++ {
		path = append(path, i)
		backtace(result, path, k-1, n, i+1)
		path = path[:len(path)-1]
	}
}
