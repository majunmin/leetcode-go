// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-23 23:10
package leetcode_0046

//https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make(map[int]struct{})
	backTrace(nums, visited, []int{}, &result)
	return result

}

func backTrace(nums []int, visited map[int]struct{}, path []int, result *[][]int) {
	// terminate
	if len(path) == len(nums) {
		dst := make([]int, len(nums))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	// for choice in choiceList
	for _, num := range nums {
		if _, exist := visited[num]; exist {
			continue
		}
		visited[num] = struct{}{}
		path = append(path, num)
		backTrace(nums, visited, path, result)
		// revert
		path = path[:len(path)-1]
		delete(visited, num)
	}
}
