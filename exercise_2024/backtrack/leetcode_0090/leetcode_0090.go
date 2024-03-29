package leetcode_0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	//param check
	// 排序
	sort.Ints(nums)
	var result [][]int
	backtrace(&result, nums, 0, []int{})
	return result
}

func backtrace(result *[][]int, nums []int, begin int, path []int) {
	// add to result
	temp := make([]int, len(path))
	copy(temp, path)
	*result = append(*result, temp)

	//for choice in choiceList
	for i := begin; i < len(nums); i++ {
		// 去掉当前层 重复的数字
		if i > begin && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		// doBacktrace
		backtrace(result, nums, i+1, path)
		path = path[:len(path)-1]
	}
}
