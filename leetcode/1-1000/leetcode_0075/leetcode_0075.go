package leetcode_0075

// https://leetcode.cn/problems/sort-colors/
func sortColors(nums []int) {
	s, e := 0, len(nums)-1
	var i int
	for i < e {
		if nums[i] == 0 {
			nums[i], nums[s] = nums[s], nums[i]
			s++
		} else if nums[i] == 2 {
			nums[i], nums[e] = nums[e], nums[i]
			e--
			i--
		}
		i++
	}
}
