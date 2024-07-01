package lcr_079

// https://leetcode.cn/problems/TVdhkn/
func subsets(nums []int) [][]int {
	var (
		result [][]int
	)
	dfs(nums, 0, []int{}, &result)
	return result
}

func dfs(nums []int, begin int, path []int, result *[][]int) {
	// termiante
	temp := make([]int, len(path))
	copy(temp, path)
	*result = append(*result, temp)
	if begin == len(nums) {
		return
	}
	for i := begin; i < len(nums); i++ {
		//path = append(path, nums[i])
		dfs(nums, i+1, append(path, nums[i]), result)
		//path = path[:len(path)-1]
	}
}
