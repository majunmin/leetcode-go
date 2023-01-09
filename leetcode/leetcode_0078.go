package leetcode

import "math"

func subsets(nums []int) [][]int {
	// param check
	if len(nums) == 0 {
		return [][]int{}
	}
	var result [][]int
	num := int(math.Pow(2.0, float64(len(nums))))
	for i := 0; i < num; i++ {
		item := i
		path := make([]int, 0, len(nums))
		idx := 0
		for item > 0 {
			if item&1 == 1 {
				path = append(path, nums[idx])
			}
			idx++
			item = item >> 1
		}

		result = append(result, path)
	}
	return result

}

func dfsSolution(nums []int) [][]int {
	var result [][]int
	dfs(&result, nums, 0, []int{})
	return result
}

func dfs(result *[][]int, nums []int, begin int, path []int) {

	dst := make([]int, len(path))
	copy(dst, path)
	*result = append(*result, dst)

	//
	for i := begin; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(result, nums, i+1, path)
		path = path[:len(path)-1]
	}

}
