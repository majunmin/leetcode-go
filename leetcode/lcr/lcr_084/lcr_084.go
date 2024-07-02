package lcr_084

// https://leetcode.cn/problems/7p8L0Z/
func permuteUnique(nums []int) [][]int {
	var (
		result  = make([][]int, 0)
		visited = make([]bool, len(nums))
	)
	backtrace(nums, visited, []int{}, &result)
	return result
}

func backtrace(nums []int, visited []bool, path []int, result *[][]int) {
	// terminate
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// for choice in choiceList
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && !visited[i-1] && nums[i] == nums[i-1] {
			continue
		}
		// 同层数字不能重复
		visited[i] = true
		path = append(path, nums[i])
		backtrace(nums, visited, path, result)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
