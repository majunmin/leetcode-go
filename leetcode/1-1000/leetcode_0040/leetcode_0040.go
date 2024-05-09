// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-01 00:55
package leetcode_0040

import "sort"

// https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i]-candidates[j] < 0
	})
	backTrace(0, []int{}, candidates, target, &result)
	return result
}

func backTrace(begin int, path []int, candidates []int, target int, result *[][]int) {
	if target == 0 {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
	}
	if target < 0 {
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 去重, 同层之间不能重复(避免结果重复)
		if i != begin && candidates[i] == candidates[i-1] {
			continue
		}
		curVal := candidates[i]
		path = append(path, curVal)
		backTrace(i+1, path, candidates, target-curVal, result)
		path = path[:len(path)-1]
	}
}
