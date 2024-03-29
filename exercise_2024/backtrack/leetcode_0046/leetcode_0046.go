package leetcode_0046

func permute(nums []int) [][]int {
	// param check
	visited := make([]bool, len(nums))
	var result [][]int
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
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true
		path = append(path, nums[i])
		backtrace(result, nums, visited, path)
		visited[nums[i]] = false
		path = path[:len(path)-1]
	}
}
