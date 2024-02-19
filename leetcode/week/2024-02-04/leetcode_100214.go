package _024_02_04

// https://leetcode.cn/contest/weekly-contest-383/problems/ant-on-the-boundary/
func returnToBoundaryCount(nums []int) int {

	// 处于边界上,   total_distance= 0
	// param check
	if len(nums) < 0 {
		// param invalid
		return 0
	}
	//
	var result int
	var total_distance int
	for i := 0; i < len(nums); i++ {
		total_distance += nums[i]
		if total_distance == 0 {
			result++
		}
	}
	return result
}
