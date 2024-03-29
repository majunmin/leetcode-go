package leetcode_0077

func combine(n int, k int) [][]int {
	// param check
	// k <= n
	var result [][]int
	backtrace(&result, n, k, 0, []int{})
	return result
}

func backtrace(result *[][]int, n int, k int, begin int, path []int) {
	// terminate
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := begin; i <= n-k; i++ {
		path = append(path, i+1)
		backtrace(result, n, k, i+1, path)
		path = path[:len(path)-1]
	}
}
