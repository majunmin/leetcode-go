package exercise_2024

// https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked
func subsets(nums []int) [][]int {
	// 位运算解法
	var result [][]int
	for i := 0; i < 1<<(len(nums)); i++ {
		path := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				path = append(path, nums[j])
			}
		}
		result = append(result, path)
	}
	return result
}

func backtraceSolution(nums []int) [][]int {
	var result [][]int
	subsetsInner(nums, 0, []int{}, &result)
	return result
}

func subsetsInner(nums []int, begin int, path []int, result *[][]int) {
	// add result
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)

	// terminate
	if begin == len(nums) {
		return
	}
	// for choice in choiceList
	for i := begin; i < len(nums); i++ {
		path = append(path, nums[i])
		subsetsInner(nums, i+1, path, result)
		path = path[:len(path)-1]
	}
}
