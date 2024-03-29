package leetcode_0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// param check
	if len(candidates) < 0 {
		panic("invalid param")
	}

	var result [][]int
	sort.Ints(candidates)
	backtrace(&result, candidates, target, 0, []int{})
	return result
}

func backtrace(result *[][]int, candidates []int, target int, begin int, path []int) {
	// terminate
	if target < 0 {
		return
	}
	if target == 0 {
		// add to result
		// return
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// for choice in choiceList
	for i := begin; i < len(candidates); i++ {
		// 同层数字不能重复
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrace(result, candidates, target-candidates[i], i+1, path) // i +1不能重复选择当前数字
		path = path[:len(path)-1]
	}
}
