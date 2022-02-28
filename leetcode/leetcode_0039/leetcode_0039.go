// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-01 00:31
package leetcode_0039

import "sort"

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	result := backTraceSolution2(candidates, target)
	return result
}

func backTraceSolution2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i]-candidates[j] < 0
	})
	backTrace2(0, []int{}, candidates, target, &result)
	return result
}

func backTrace2(begin int, path []int, candidates []int, target int, result *[][]int) {
	if target == 0 {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}
	if target < 0 {
		return
	}

	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backTrace2(i, path, candidates, target-candidates[i], result)
		path = path[:len(path)-1]
	}
}

func backTraceSolution1(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	// 升序排序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i]-candidates[j] < 0
	})
	backTrace(0, []int{}, 0, candidates, target, &result)
	return result
}

// backTrace
// idx: 当前索引
// path: 已选择的组合
// sum: 当前 path 的总和
func backTrace(begin int, path []int, sum int, candidates []int, target int, result *[][]int) {
	// terminate
	if sum == target {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	if sum > target {
		return
	}

	// for choice in choiceList
	for i := begin; i < len(candidates); i++ {
		// do choice
		path = append(path, candidates[i])
		// process
		backTrace(i, path, sum+candidates[i], candidates, target, result)
		// revert choice
		path = path[:len(path)-1]
	}
}
