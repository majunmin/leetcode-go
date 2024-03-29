package leetcode_0078

// 1. 回溯算法 dfs
// 2. 位运算
// 3. 递归解法
func subsets2(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		panic("invalid param")
	}
	//
	cnt := 1 << len(nums)
	var result [][]int
	for i := 0; i < cnt; i++ {
		path := make([]int, 0, len(nums))
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				path = append(path, nums[j])
			}
		}
		result = append(result, path)
	}
	return result
}

func subsets3(nums []int) [][]int {
	var result [][]int
	dfs(&result, nums, 0, []int{})
	return result
}

func dfs(result *[][]int, candidates []int, cur int, path []int) {
	// terminate
	if cur == len(candidates) {
		// copy path to result
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// choose this num
	path = append(path, candidates[cur])
	dfs(result, candidates, cur+1, path)
	// not choose this num
	path = path[:len(path)-1]
	dfs(result, candidates, cur+1, path)
}

func subsets1(nums []int) [][]int {
	var result [][]int
	backtrace(&result, nums, 0, []int{})
	return result
}

func backtrace(result *[][]int, candidates []int, begin int, path []int) {
	// add result
	temp := make([]int, len(path))
	copy(temp, path)
	*result = append(*result, temp)
	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtrace(result, candidates, i+1, path)
		path = path[:len(path)-1]
	}
}
