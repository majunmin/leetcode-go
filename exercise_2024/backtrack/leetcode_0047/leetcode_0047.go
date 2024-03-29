package leetcode_0047

import "sort"

func permuteUnique(nums []int) [][]int {
	// param check
	sort.Ints(nums)
	var result [][]int
	visited := make([]bool, len(nums))
	backtrace(&result, nums, visited, []int{})
	return result
}

func backtrace(result *[][]int, nums []int, visited []bool, path []int) {
	// terminate
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// for choice in choiceList
	for i := 0; i < len(nums); i++ {
		if visited[i] || (i > 0 && nums[i] == nums[i-1] && visited[i-1]) {
			continue
		}
		visited[i] = true
		path = append(path, nums[i])
		backtrace(result, nums, visited, path)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
