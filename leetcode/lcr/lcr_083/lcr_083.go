package lcr_083

// https://leetcode.cn/problems/VvJkup/
func permute(nums []int) [][]int {
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
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		path = append(path, nums[i])
		backtrace(nums, visited, path, result)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
