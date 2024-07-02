package lcr_082

import "sort"

// https://leetcode.cn/problems/4sjJUc/
func combinationSum2(candidates []int, target int) [][]int {
	var (
		result = make([][]int, 0)
	)
	sort.Ints(candidates)
	backtrace(candidates, 0, target, []int{}, &result)
	return result
}

func backtrace(candidates []int, begin int, target int, path []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := begin; i < len(candidates); i++ {
		// 同层也不能重复
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrace(candidates, i+1, target, path, result)
		path = path[:len(path)-1]
	}
}
