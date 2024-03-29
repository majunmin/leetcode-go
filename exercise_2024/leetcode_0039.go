package exercise_2024

// https://leetcode.cn/problems/combination-sum/?envType=study-plan-v2&envId=top-100-liked
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	backtrace39(candidates, target, 0, &result, []int{})
	return result
}

func backtrace39(candidates []int, target int, begin int, result *[][]int, path []int) {
	// terminate
	if target < 0 {
		// 剪枝
		return
	}
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// process
	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtrace39(candidates, target-candidates[i], i, result, path)
		path = path[:len(path)-1]
	}
}
