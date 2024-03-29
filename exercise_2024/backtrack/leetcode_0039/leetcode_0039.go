package leetcode_0039

func combinationSum(candidates []int, target int) [][]int {
	// param check
	if len(candidates) == 0 {
		panic("invalid param")
	}
	var result [][]int
	backtrace(&result, candidates, target, 0, []int{})
	return result
}

func backtrace(result *[][]int, candidates []int, target int, begin int, path []int) {
	// terminate
	// 剪枝
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// for choice in choiceList
	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtrace(result, candidates, target-candidates[i], i, path) // 同一个数 可以继续选择,所以这里传 i,而不是 i + 1
		path = path[:len(path)-1]
	}
}
