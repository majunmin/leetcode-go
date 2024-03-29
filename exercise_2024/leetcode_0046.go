package exercise_2024

// https://leetcode.cn/problems/permutations/?envType=study-plan-v2&envId=top-100-liked
func permute(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	//
	var result [][]int
	visited := make([]bool, len(nums))
	permuteInner(nums, []int{}, visited, &result)
	return result
}

func permuteInner(nums []int, path []int, visited []bool, result *[][]int) {
	// terminate
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// for choice in choiceList
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		path = append(path, nums[i])
		permuteInner(nums, path, visited, result)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
