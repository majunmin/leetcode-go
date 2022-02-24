// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-24 23:19
package leetcode_0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	//先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backTrace(0, []int{}, nums, &result)
	return result
}

func backTrace(begin int, path []int, nums []int, result *[][]int) {
	// add
	dst := make([]int, len(path))
	copy(dst, path)
	*result = append(*result, dst)

	// for choice in choiceList
	for i := begin; i < len(nums); i++ {
		// 剪枝, 同层 不选相同的数字
		if i != begin && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backTrace(i+1, path, nums, result)
		// revert
		path = path[:len(path)-1]
	}
}
