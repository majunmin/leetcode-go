package exercise_2024

// https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	var result [][]int
	visited := make([]bool, len(nums))
	permuteUniqueInner(nums, visited, []int{}, &result)
	return result
}

func permuteUniqueInner(nums []int, visited []bool, path []int, result *[][]int) {
	// terminate
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// for choice in choiceList
	for i, num := range nums {
		if visited[i] {
			continue
		}
		// 同层相同数字去重(剪枝)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		visited[i] = true
		path = append(path, num)
		permuteUniqueInner(nums, visited, path, result)
		visited[i] = false
		path = path[:len(path)-1]
	}
}
