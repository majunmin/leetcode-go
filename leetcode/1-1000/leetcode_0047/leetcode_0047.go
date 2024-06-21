// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-24 00:09
package leetcode_0047

import "sort"

// https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	// store index that has been visited
	visited := make(map[int]bool)
	result := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backTrace([]int{}, visited, nums, &result)
	return result
}

// 回溯算法
func backTrace(path []int, visited map[int]bool, nums []int, result *[][]int) {
	if len(path) == len(nums) {
		// copy path
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
	}

	for i, num := range nums {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		path = append(path, num)
		visited[i] = true
		backTrace(path, visited, nums, result)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
