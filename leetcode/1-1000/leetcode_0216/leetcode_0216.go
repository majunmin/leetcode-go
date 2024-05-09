package leetcode_0216

// https://leetcode.cn/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	// param check
	var (
		result = make([][]int, 0)
	)
	backtrace(n, k, 1, []int{}, &result)
	return result
}

func backtrace(target int, level int, begin int, path []int, result *[][]int) {
	// terminate
	if target < 0 {
		return
	}
	if target == 0 && level == 0 { //要求长度为k， 所以这里限制下  level == 0
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	if level == 0 {
		return
	}
	for i := begin; i < 10; i++ {
		path = append(path, i)
		backtrace(target-i, level-1, i+1, path, result)
		path = path[:len(path)-1]
	}
}
