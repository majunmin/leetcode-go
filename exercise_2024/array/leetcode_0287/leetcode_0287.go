package leetcode_0287

// https://leetcode.cn/problems/find-the-duplicate-number/?envType=study-plan-v2&envId=top-100-liked
func findDuplicate(nums []int) int {
	//
	if len(nums) < 2 {
		//panic("invalid param")
		return -1
	}
	visited := make(map[int]bool)
	item := nums[0]
	for !visited[item] {
		visited[item] = true
		item = nums[item]
	}
	return item
}
