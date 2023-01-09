package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	dfs2(&result, nums, 0, []int{})
	return result
}

func dfs2(result *[][]int, nums []int, begin int, path []int) {

	dst := make([]int, len(path))
	copy(dst, path)
	*result = append(*result, dst)

	// same level can not cantains  same num
	for i := begin; i < len(nums); i++ {
		// 剪枝
		if i != begin && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs2(result, nums, i+1, path)
		path = path[:len(path)-1]
	}
}
