package exercise_2024

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// param check
	if len(candidates) < 0 {
		return nil
	}
	var result [][]int
	sort.Ints(candidates)
	backtrace40(candidates, target, 0, &result, []int{})
	return result
}

func backtrace40(candidates []int, target int, begin int, result *[][]int, path []int) {
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

	for i := begin; i < len(candidates); i++ {
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrace40(candidates, target-candidates[i], i+1, result, path)
		path = path[:len(path)-1]
	}
}
