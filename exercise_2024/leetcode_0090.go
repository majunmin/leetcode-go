package exercise_2024

// https://leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	//
	var result [][]int
	subsetsWithDupInner(nums, 0, []int{}, &result)
	return result
}

func subsetsWithDupInner(nums []int, begin int, path []int, result *[][]int) {
	// add result
	tmp := make([]int, len(nums))
	copy(tmp, path)
	*result = append(*result, tmp)
	// terminate
	if begin == len(nums) {
		return
	}

	// for choice in choiceList
	for i := begin; i < len(nums); i++ {
		if i > begin && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		subsetsWithDupInner(nums, i+1, path, result)
		path = path[:len(path)-1]
	}
}
